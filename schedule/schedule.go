package schedule

import(
	"log"
	"os"
	// "sync"
	"fmt"
	// "errors"
	"time"
	"strings"
	"io/ioutil"
	"path/filepath"
	"github.com/robfig/cron/v3"
	"github.com/ruvido/letter/markdown"
	"github.com/ruvido/letter/send"
	"github.com/spf13/viper"
)

func Send() {

	searchScheduledLetters()
    // destroyLock() //clean leftovers from previous sessions

	// start cron
	c := cron.New()
	crontab := viper.GetString("schedule.crontab")
	c.AddFunc(crontab, searchScheduledLetters )
	c.Start()

	select {} // Keep the program running indefinitely
}

func searchScheduledLetters() {

    // createLock() // Lock for the entire function
    // defer destroyLock() // Unlock at the end of the function

	content := viper.GetString("schedule.content")
	archive := viper.GetString("schedule.archive") 

	emails := listEmails(content)
    for _, em := range emails {
        // log.Println(em.Date, em.Filename)
        collectionName   := em.Schedule.Collection
        collectionFilter := em.Schedule.Filter
        if collectionName == "" {
            collectionName = viper.GetString("schedule.collection")
        }
        if collectionFilter == "" {
            collectionFilter = viper.GetString("schedule.filter")
        }

        if true {
            if err := send.Newsletter(em.Filename,collectionName,collectionFilter); err != nil {
                log.Fatalf("Error newsletter sending: %v", err)
            }

            if err := archiveEmail(em, archive); err != nil {
                log.Fatalf("Error archiving email: %v", err)
            } else {
                fmt.Printf("Successfully archived file: %s\n", em.Subject)
            }
        } else {
            // DEBUGGING
            log.Println("Sending Sleeping test of 90s")
            time.Sleep(90 * time.Second)
            log.Println("Sleep ends")
        }
	}
}

func listEmails(folder string) []markdown.Email {

	list        := []markdown.Email{} // Initializes an empty slice
	futureList  := []markdown.Email{}
	
	// Read the directory
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over each file in the directory
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			filePath := filepath.Join(folder, file.Name()) // Use "content" directory here
			// Read the file content
			email, err := markdown.BuildEmail(filePath)
			if err != nil {
				log.Printf("Failed to build email from file: %s, error: %v\n", filePath, err)
				continue
			}

			today := time.Now().Format("2006-01-02")
			edate := email.Date.Format("2006-01-02")

			if edate == today {
				list = append(list, email)
			}
			if edate > today {
				futureList = append(futureList, email)
            }
		}
	}

    // Check if the list is empty and print the appropriate message
    if len(futureList) > 0 {
        fmt.Println("")
        fmt.Println("FUTURE SCHEDULES")
        for _, email := range futureList {
            fmt.Printf("  * %s - %s | %s \n", 
            email.Date.Format("2006-01-02"), email.Subject, email.Schedule.Collection)
        }
    }
    fmt.Println("")
    fmt.Println("TODAYS SCHEDULES")
    if len(list) == 0 {
        fmt.Println("No letters are scheduled for today")
    } else {
        for _, email := range list {
            fmt.Printf("  * %s - %s | %s \n", 
            email.Date.Format("2006-01-02"), email.Subject, email.Schedule.Collection)
        }
    }
    fmt.Println("")

	return list
}

func archiveEmail(em markdown.Email, archive string) error {
	// Get the source file path
	sourcePath := em.Filename
	archiveFolder := archive
	
	// Ensure the archive directory exists
	if _, err := os.Stat(archiveFolder); os.IsNotExist(err) {
		log.Println("archive folder does not exist")
		if err := os.MkdirAll(archiveFolder, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create archive directory: %v", err)
		}
	}

	// Get the destination file path
	destPath := filepath.Join(archiveFolder, filepath.Base(sourcePath))
	
	// Move the file to the archive directory
	if err := os.Rename(sourcePath, destPath); err != nil {
		return fmt.Errorf("failed to move file from %s to %s: %v", sourcePath, destPath, err)
	}

	return nil
}

// // Locking mechanism
// var (
//     mu            sync.Mutex // Create a mutex lock
//     isExecuting   bool       // Flag to indicate if the function is currently executing
// )
//
// func createLock() {
//     mu.Lock() // Acquire the lock
//     log.Println("Session lock")
//     if isExecuting {
//         log.Println("Alert: Attempt to execute searchScheduledLetters while it is already running.")
//     }
//     isExecuting = true // Set the flag to indicate execution
// }
//
// func destroyLock() {
//     isExecuting = false // Reset the flag
//     mu.Unlock() // Release the lock
//     log.Println("LOCK DESTROYED")
// }

package schedule

import(
	"log"
	"os"
	"fmt"
	// "errors"
	"time"
	"strings"
	"io/ioutil"
	"path/filepath"
	"github.com/robfig/cron/v3"
	"github.com/ruvido/letter/markdown"
	"github.com/spf13/viper"
	
)
	
func Send() {

	log.Println("Looking for letters...")
	crontab := viper.GetString("schedule.crontab")
	c := cron.New()
	c.AddFunc(crontab, searchScheduledLetters )
	c.Start()

	select {} // Keep the program running indefinitely
}

func searchScheduledLetters() {

	content := viper.GetString("schedule.content")
	archive := viper.GetString("schedule.archive") 

	emails := listEmails(content)
	for _, em := range emails {
		log.Println(em.Date, em.Filename)
		if err := archiveEmail(em, archive); err != nil {
			log.Fatalf("Error archiving email: %v", err)
		} else {
			fmt.Printf("Successfully archived file: %s\n", em.Filename)
		}
	}

}

func listEmails(folder string) []markdown.Email {

	list := []markdown.Email{} // Initializes an empty slice
	
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
				// log.Println(email.Date)
				// log.Printf("Sending file: %s\n", filePath)
				list = append(list, email)
			}
		}
	}

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

package schedule

import(
	"log"
	"os"
	// "errors"
	"time"
	"strings"
	"io/ioutil"
	"path/filepath"
	"github.com/robfig/cron/v3"
	"github.com/ruvido/letter/markdown"
)
	
func Send() {
	log.Println("sbam")
	searchScheduledLetters()
	c := cron.New()
	c.AddFunc("* * * * *", searchScheduledLetters )
	c.Start()

	select {} // Keep the program running indefinitely
}

func searchScheduledLetters() {
	log.Println("boombasticsaydai")

	files := listFiles("content")

	// Iterate over each file in the directory
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			filePath := filepath.Join("content", file.Name()) // Use "content" directory here

			// Read the file content
			email, err := markdown.BuildEmail(filePath)
			if err != nil {
				log.Printf("Failed to build email from file: %s, error: %v\n", filePath, err)
				continue
			}

			// Check if the file contains YAML front matter
			// Get today's date in YYYY-MM-DD format
			today := time.Now().Format("2006-01-02")
			edate := email.Date.Format("2006-01-02")

			if edate == today {
				log.Println(email.Date)
				log.Printf("Sending file: %s\n", filePath)
			}
		}
	}
}

func listFiles(folder string) []os.FileInfo {
	// Read the directory
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

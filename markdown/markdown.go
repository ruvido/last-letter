package markdown
                                
import (
	"log"
    "os"
    "bytes"
    "time"
    "strings"
    "io/ioutil"
    "github.com/yuin/goldmark"
    "github.com/adrg/frontmatter"
)

type Email struct {
    Date     time.Time `yaml:"date"`
    Subject  string    `yaml:"title"`
	Tags  []string     `yaml:"tags"`
    Content  string
    Txt      string
}

func BuildEmail ( markdownFilename string ) Email {

	var em Email
	
	// Read whole markdown file
	wholeFile, err := ioutil.ReadFile(markdownFilename)
	if err != nil {
		log.Println("Error: loading markdown file "+markdownFilename)
	}

	// Extract frontmatter + body text
	body, err := frontmatter.Parse(strings.NewReader(string(wholeFile)), &em)
    if err != nil {
        log.Println("Error: reading the front matter")
    }

	// Abort if Subject is missing
	if em.Subject == "" {
		log.Println("Error: subject is missing")
		os.Exit(99)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(body, &buf); err != nil {
		panic(err)
	}

	em.Content = buf.String()

	return em
}

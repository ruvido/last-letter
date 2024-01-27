package markdown
                                
import (
	"log"
    "os"
    "bytes"
    "time"
    "strings"
    "io/ioutil"
	"html/template"
    "github.com/yuin/goldmark"
    "github.com/adrg/frontmatter"
	"github.com/spf13/viper"
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

	// em.Content = buf.String()
	em.Content = EmailTemplating(buf.String())

	return em
}

func EmailTemplating ( htmlContent string ) string {
	// Read the template from a file
	templateHtml := viper.GetString ("general.template")+".html"
	t, err := template.ParseFiles(templateHtml)
	if err != nil {
		panic(err)
	}

	// Create a data structure to hold your content
	data := struct {
		// Title   string
		// Heading string
		Content template.HTML
	}{
		// Title:   "My Page",
		// Heading: "Welcome to My Page",
		Content: template.HTML(htmlContent),
	}

	// Create a buffer to hold the filled template
	var tplBuffer bytes.Buffer

	// Write the filled template to the buffer
	err = t.Execute(&tplBuffer, data)
	if err != nil {
		panic(err)
	}

	// Convert the buffer to a string
	filledTemplate := tplBuffer.String()

	return filledTemplate
}

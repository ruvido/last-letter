package markdown
                                
import (
    "bytes"
    "time"
	"errors"
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
    Filename string
    Schedule struct {
        Collection  string `yaml:"collection"`
        Filter      string `yaml:"filter"`
    } `yaml:"schedule"`
}

func BuildEmail ( markdownFilename string ) (Email, error) {

	var em Email
	
	em.Filename = markdownFilename
	
	// Read whole markdown file
	wholeFile, err := ioutil.ReadFile(markdownFilename)
	if err != nil {
		return em, err
	}

	// Extract frontmatter + body text
	body, err := frontmatter.Parse(strings.NewReader(string(wholeFile)), &em)
	if err != nil {
		return em, err
	}
	
	// Abort if Subject is missing
	if em.Subject == "" {
		return em, errors.New("subject is missing")
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(body, &buf); err != nil {
		return em, err
	}

	// Set the email content
	em.Content = EmailTemplating(buf.String())

	return em, nil
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

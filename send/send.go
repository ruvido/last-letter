package send

import(
	"fmt"
	"log"
	"time"
	"github.com/ruvido/letter/pkg"
	"github.com/ruvido/letter/markdown"
	gomail "gopkg.in/gomail.v2"
	"github.com/spf13/viper"
    // "github.com/gosimple/slug"
	// "os"

)


func Newsletter( markdownFile string, collection string, filter string ){

	// Fetch email addresses from database collection
	emailAddrs := letter.GetEmailsFrom( collection, filter )

	// log.Println(emailAddrs)
	// os.Exit(99)

	// Create email object from markdown file
	em := markdown.BuildEmail( markdownFile )


	// Batch send emails
	smtpSend( em.Subject, em.Content, em.Txt, emailAddrs ) 


}

func smtpSend ( subject string, content string, txt string, addrs []string ){

	user := viper.GetString ("smtp.user")
	pass := viper.GetString ("smtp.password")
	addr := viper.GetString ("smtp.address")
	port := viper.GetInt    ("smtp.port")
	sendr:= viper.GetString ("general.sender")

	gm := gomail.NewDialer(addr, port, user, pass)
    s, err := gm.Dial()
    // _, err := gm.Dial()
    if err != nil {panic(err)}

	m := gomail.NewMessage()

	log.Println(len(addrs))
	batch := viper.GetInt("sending.batch")
	waits := viper.GetInt("sending.waits")
	cc := 0
	for i := 0; i < len(addrs); i += batch {
		end := i + batch
		if end > len(addrs) {
			end = len(addrs)
		}
		slice := addrs[i:end]

		if cc>0 {time.Sleep(time.Duration(waits) * time.Second)}
		for _, addr := range slice {
			cc+=1
			m.SetHeader ("From", sendr )
			m.SetHeader ("To",        addr )
			m.SetHeader ("Subject",   subject )
			m.SetBody   ("text/html", content )
			fmt.Printf("SEND> %4d/%-4d   %s\n", cc,len(addrs),addr)
			if err := gomail.Send(s, m); err != nil {
				log.Printf("Could not send email to %q: %v", addr, err)
			}
			m.Reset()

		}
	}
}

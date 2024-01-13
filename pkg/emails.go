package letter

import(
	"log"

)

func SendingNewsletter( markdown string, collection string ){
	emails := GetEmailsFrom( collection )
	log.Println(len(emails))
	log.Println("letter.MarkdownToHtml( mardown )     >> convert markdown")
	log.Println("letter.SmtpSendsend( html, emails )  >> smtp send")
}

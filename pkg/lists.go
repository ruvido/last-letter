package letter

import (
    "github.com/spf13/viper"
	"log"
	"fmt"
)
func GetEmailsFrom( collection string ) []string {
	var emails []string
	if collection != "" {
		log.Println("get emails from collection -->", collection)
		emails = PocketbaseEmailsFrom( collection )

	} else {
		// emails = []string{viper.GetString("test.email")}
		emails = append(emails, viper.GetString("test.email"))
		log.Println("send a test to",emails)
	}
	return emails
}

func DumpArray( array []string ) {
	for idx,item := range array {
		fmt.Println(idx,item)
	}

}

package letter

import (
	"github.com/ruvido/letter/globals" // import our globals package
    "github.com/spf13/viper"
	"log"
	"fmt"
    "strings"
	// "os"
)
func GetEmailsFrom( collection, filter string ) []string {
	var emails []string
	if collection != "" {
		log.Println("get emails from collection -->", collection)
		emails = PocketbaseEmailsFrom( collection, filter )
    } else {
        if globals.AltList != "" {
			emails = strings.Split(globals.AltList, ",")
            // os.Exit(99)
        } else {
            emails = append(emails, viper.GetString("test.email"))
            log.Println("send a test to",emails)
        }
    }
	return emails
}

func DumpArray( array []string ) {
	for idx,item := range array {
		fmt.Println(idx,item)
	}

}

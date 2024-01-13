package letter

import (
	"log"
	"github.com/spf13/viper"
	// "github.com/r--w/pocketbase"
	"github.com/pluja/pocketbase"

)

func PocketbaseEmailsFrom( collection string ) []string {

	client := pocketbase.NewClient(
		viper.GetString("pocketbase.address"),
		pocketbase.WithAdminEmailPassword( 
			viper.GetString("pocketbase.admin"),
			viper.GetString("pocketbase.password")))

	response, err := client.List(
		collection, pocketbase.ParamsList{
        Page: 1, Size: 10000, Sort: "-created", 
		Filters: viper.GetString("pocketbase.filters"),
    })
    if err != nil {
        log.Fatal(err)
    }
    // log.Printf("Total of Emails: %d\n",response.TotalItems)

	var emails []string
	for _, item := range response.Items {
		emails = append(emails, item["email"].(string))
	}

	return emails
}

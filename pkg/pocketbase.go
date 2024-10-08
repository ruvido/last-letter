package letter

import (
	"log"
	"github.com/spf13/viper"
	// "github.com/r--w/pocketbase"
	"github.com/pluja/pocketbase"

)

// func PocketbaseUpdateRecord( collection string ) {
//
// }

func PocketbaseEmailsFrom( collection string, filter string ) []string {

	client := pocketbase.NewClient(
		viper.GetString("pocketbase.address"),
		pocketbase.WithAdminEmailPassword( 
			viper.GetString("pocketbase.admin"),
			viper.GetString("pocketbase.password")))

	
	// log.Println(collection)
	// os.Exit(5)
	emails := []string{}
	keepListing := true
	askPage := 1
	for keepListing {
		response, err := client.List(
			collection, pocketbase.ParamsList{
				Page: askPage, Size: 500, Sort: "-created", 
				// Filters: viper.GetString("pocketbase.filters"),
				Filters: filter,
				// Page: 1, Size: 10000, Sort: "-created", 
		})
		if err != nil {
			log.Println("bum")
			log.Fatal(err)
		}

		// fmt.Println("page: ",response.Page)
		// fmt.Println("perpage:",response.PerPage, len(response.Items))
		// fmt.Println("totalitems:",response.TotalItems)
		// fmt.Println("totalpages:",response.TotalPages)

		page := response.Page
		totp := response.TotalPages

		if page < totp {
			askPage = page+1
		} else {
			keepListing = false
		}

		for _, item := range response.Items {
			emails = append(emails, item["email"].(string))
		}
	}

	return emails
}

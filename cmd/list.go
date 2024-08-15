package letter

import (
	"github.com/ruvido/letter/pkg"
    "github.com/spf13/cobra"
    "fmt"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "Print email addresses for list",
	// Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		// if len(args) > 0 {
		if len(args) >= 0 {
			fmt.Println("Emails addresses from list")
			// emails := letter.GetEmailsFrom(args[0], collectionFilter)
			emails := letter.GetEmailsFrom(collectionName, collectionFilter)
			letter.DumpArray(emails)
		} else {
			fmt.Println("\nError> Missing collection name\n")
		}
	},
}

func init() {
    rootCmd.AddCommand(listCmd)
}

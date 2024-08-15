package letter

import (
    // "github.com/spf13/viper"
    "github.com/spf13/cobra"
    "github.com/ruvido/letter/send"
    "fmt"
	// "os"
)

var (
	collectionFlag string // email addresses collection
)

var sendCmd = &cobra.Command{
    Use:   "send",
    Short: "Send a newsletter to collection",
	// Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			// send.Newsletter(args[0],collectionFlag)
			send.Newsletter(args[0],collectionName,collectionFilter)
		} else {
			fmt.Println("\nletter send <markdown.file> (-c <collection.name> -f <collection.filter>)")
			fmt.Println("Error> Missing markdown file\n")
		}
    },
}

func init() {
	// sendCmd.Flags().StringVarP(&collectionFlag, "collection", "c", "", "email addresses collection name")
    rootCmd.AddCommand(sendCmd)
}

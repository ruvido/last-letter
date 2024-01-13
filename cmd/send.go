package letter

import (
    // "github.com/spf13/viper"
    "github.com/spf13/cobra"
    "github.com/ruvido/letter/pkg"
    "fmt"
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
			letter.SendingNewsletter(args[0],collectionFlag)
		} else {
			fmt.Println("\nError> Missing markdown file\n")
		}
    },
}

func init() {
	sendCmd.Flags().StringVarP(&collectionFlag, "collection", "c", "", "email addresses collection name")
    rootCmd.AddCommand(sendCmd)
}

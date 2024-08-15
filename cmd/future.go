package letter

import (
    "github.com/spf13/cobra"
    "fmt"
)

var futureCmd = &cobra.Command{
    Use:   "future",
    Short: "Show me future scheduled campaigns",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Showing future campaigns")
    },
}

func init() {
    rootCmd.AddCommand(futureCmd)
}

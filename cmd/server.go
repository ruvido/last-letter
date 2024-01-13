package letter

import (
    "github.com/spf13/cobra"
    "fmt"
)

var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Newsletter server sending scheduled campaigns to list",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Newsletter server running...")
    },
}

func init() {
    rootCmd.AddCommand(serverCmd)
}

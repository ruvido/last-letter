package letter

import (
    "github.com/spf13/cobra"
    "github.com/ruvido/letter/server"
	// "fmt"
)

var serverCmd = &cobra.Command{
    Use:   "serve",
    Short: "Newsletter server listening to scheduled campaigns",
    Run: func(cmd *cobra.Command, args []string) {
		server.Start()
    },
}

func init() {
    rootCmd.AddCommand(serverCmd)
}

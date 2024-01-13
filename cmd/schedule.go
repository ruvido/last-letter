package letter

import (
	"github.com/ruvido/letter/pkg"
    "github.com/spf13/cobra"
    "fmt"
)

var scheduleCmd = &cobra.Command{
    Use:   "schedule",
    Short: "Schedule campaigns in file",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Scheduling campaigns in file")
		letter.Schedule(args[0])
    },
}

func init() {
    rootCmd.AddCommand(scheduleCmd)
}

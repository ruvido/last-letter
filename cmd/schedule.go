package letter

import (
	"github.com/ruvido/letter/schedule"
    "github.com/spf13/cobra"
    "log"
)

var scheduleCmd = &cobra.Command{
    Use:   "schedule",
    Short: "Schedule campaigns in file",
    Run: func(cmd *cobra.Command, args []string) {
        log.Println("Letter scheduler started")
		//letter.Schedule(args[0])
		schedule.Send()
    },
}

func init() {
    rootCmd.AddCommand(scheduleCmd)
}

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
        log.Println("Scheduling campaigns in file")
		//letter.Schedule(args[0])
		schedule.Send()
    },
}

func init() {
    rootCmd.AddCommand(scheduleCmd)
}

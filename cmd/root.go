package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rscheduler-cli/cmd/infoCMD"
	"rscheduler-cli/cmd/processorCMD"
	"rscheduler-cli/cmd/rmpCMD"
	"rscheduler-cli/cmd/rmtCMD"
	"rscheduler-cli/cmd/taskCMD"
)

var rootCmd = &cobra.Command{
	Use:   "rscheduler",
	Short: "rscheduler cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello rscheduler!")
		fmt.Println("Try using --help for more help information")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(
		infoCMD.InfoCmd,
		processorCMD.ProcessorCMD,
		taskCMD.TaskCMD,
		rmpCMD.RmpCMD,
		rmtCMD.RmtCMD,
	)
}

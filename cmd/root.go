package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"rscheduler-cli/cmd/infoCMD"
)

var rootCmd = &cobra.Command{
	Use:   "rscheduler",
	Short: "rscheduler cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello rscheduler!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(infoCMD.InfoCmd)
}

package infoCMD

import "fmt"

var infoVersionFlag bool

func init() {
	InfoCmd.Flags().BoolVarP(&infoVersionFlag, "version", "v", false, "Get rscheduler version")
}

func getInfoVersion() string {
	bi := getBaseInfo()
	return fmt.Sprintf("rscheduler version: %s", bi.Version)
}

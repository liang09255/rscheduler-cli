package infoCMD

import "fmt"

var infoVersionFlag bool

func init() {
	InfoCmd.Flags().BoolVarP(&infoVersionFlag, "version", "v", false, "Get rscheduler version")
}

func getInfoVersion() (string, error) {
	bi, err := getBaseInfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("rscheduler version: %s", bi.Version), nil
}

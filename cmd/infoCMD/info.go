package infoCMD

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"rscheduler-cli/config"
	"rscheduler-cli/pkg/model"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "rscheduler info",
	RunE: func(cmd *cobra.Command, args []string) error {
		var str string
		var err error
		if infoVersionFlag {
			str, err = getInfoVersion()
			if err != nil {
				return err
			}
		} else {
			str, err = getInfo()
			if err != nil {
				return err
			}
		}
		fmt.Println(str)
		return nil
	},
}

func getBaseInfo() (*model.BaseInfo, error) {
	request := gorequest.New()
	_, body, errs := request.Get(config.Config.BaseURL + "/base/info").End()
	if errs != nil {
		return nil, errs[0]
	}
	var baseInfoResp model.BaseInfoResp
	err := json.Unmarshal([]byte(body), &baseInfoResp)
	if err != nil {
		return nil, err
	}
	return &baseInfoResp.BaseInfo, nil
}

func getInfo() (string, error) {
	bi, err := getBaseInfo()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%+v", bi), nil
}

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
	Run: func(cmd *cobra.Command, args []string) {
		var str string
		if infoVersionFlag {
			str = getInfoVersion()
		} else {
			str = getInfo()
		}
		fmt.Println(str)
	},
}

func getBaseInfo() model.BaseInfo {
	request := gorequest.New()
	_, body, errs := request.Get(config.Config.BaseURL + "/base/info").End()
	if errs != nil {
		fmt.Println("get version failed, err: " + errs[0].Error())
	}
	var baseInfoResp model.BaseInfoResp
	err := json.Unmarshal([]byte(body), &baseInfoResp)
	if err != nil {
		fmt.Println("unmarshal version failed, err: " + err.Error())
	}
	return baseInfoResp.BaseInfo
}

func getInfo() string {
	bi := getBaseInfo()
	return fmt.Sprintf("%+v", bi)
}

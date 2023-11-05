package rmtCMD

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"rscheduler-cli/config"
	"rscheduler-cli/pkg/model"
)

var RmtCMD = &cobra.Command{
	Use:   "rmt",
	Short: "remove task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("please provide taskID")
		}
		return RemoveTask(args[0])
	},
}

func RemoveTask(taskID string) error {
	request := gorequest.New()
	_, body, errs := request.Post(config.Config.BaseURL+"/task/delete").
		Param("id", taskID).
		End()
	if errs != nil {
		return errs[0]
	}
	var deleteResp model.TaskDeleteResp
	err := json.Unmarshal([]byte(body), &deleteResp)
	if err != nil {
		return err
	}
	if deleteResp.Success {
		fmt.Println("success")
		return nil
	}
	return fmt.Errorf("fail: %s", deleteResp.Info)
}

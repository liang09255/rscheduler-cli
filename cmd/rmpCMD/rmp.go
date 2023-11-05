package rmpCMD

import (
	"encoding/json"
	"fmt"
	"github.com/liang09255/lutils/conv"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"rscheduler-cli/config"
	"rscheduler-cli/pkg/model"
)

var RmpCMD = &cobra.Command{
	Use:   "rmp",
	Short: "remove processor",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("please provide processorID")
		}
		return RemoveProcessor(args[0])
	},
}

func RemoveProcessor(processorID string) error {
	request := gorequest.New()
	_, body, errs := request.Post(config.Config.BaseURL+"/processor/delete").
		Param("id", processorID).
		Param("force", conv.ToString(force)).
		End()
	if errs != nil {
		return errs[0]
	}
	var deleteResp model.ProcessorDeleteResp
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

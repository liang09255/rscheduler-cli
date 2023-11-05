package processorCMD

import (
	"encoding/json"
	"github.com/liang09255/lutils/conv"
	"github.com/olekukonko/tablewriter"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"os"
	"rscheduler-cli/config"
	"rscheduler-cli/pkg/model"
	"strconv"
)

var ProcessorCMD = &cobra.Command{
	Use:   "processor",
	Short: "processor info",
	RunE: func(cmd *cobra.Command, args []string) error {
		table, err := getProcessorInfoTable()
		if err != nil {
			return err
		}
		table.Render()
		return nil
	},
}

func getProcessorInfoTable() (*tablewriter.Table, error) {
	table := tablewriter.NewWriter(os.Stdout)
	headers := []string{
		"ID", "Name", "PID", "TotalTask", "Running", "TaskID",
	}
	table.SetHeader(headers)
	infos, err := getProcessorInfo()
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		strList := []string{
			info.ID,
			info.Name,
			strconv.Itoa(info.PID),
			strconv.Itoa(int(info.TotalTaskNum)),
			conv.ToString(info.Running),
			info.TaskID,
		}
		table.Append(strList)
	}
	return table, nil
}

func getProcessorInfo() ([]model.ProcessorInfo, error) {
	request := gorequest.New()
	_, body, errs := request.Get(config.Config.BaseURL + "/processor/info").End()
	if errs != nil {
		return nil, errs[0]
	}
	var processorInfoResp model.ProcessorInfoResp
	err := json.Unmarshal([]byte(body), &processorInfoResp)
	if err != nil {
		return nil, err
	}
	return processorInfoResp.ProcessorInfo, nil
}

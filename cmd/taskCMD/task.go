package taskCMD

import (
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"os"
	"rscheduler-cli/config"
	"rscheduler-cli/pkg/model"
	"strconv"
	"time"
)

var TaskCMD = &cobra.Command{
	Use:   "task",
	Short: "task info",
	RunE: func(cmd *cobra.Command, args []string) error {
		table, err := getTaskInfoTable()
		if err != nil {
			return err
		}
		table.Render()
		return nil
	},
}

func getTaskInfoTable() (*tablewriter.Table, error) {
	table := tablewriter.NewWriter(os.Stdout)
	headers := []string{
		"ID", "Name", "ProcessorID", "Memory", "CPU", "CreateAt", "StartAt", "RunTime",
	}
	table.SetHeader(headers)

	taskInfos, err := getTaskInfo()
	if err != nil {
		return nil, err
	}

	memoryFormat := func(num uint64) string {
		return strconv.FormatUint(num, 10) + "MB"
	}
	cpuFormat := func(num uint64) string {
		return strconv.FormatUint(num, 10) + "%"
	}
	timeFormat := func(t int64) string {
		return time.Unix(t, 0).Format(time.DateTime)
	}
	secondFormat := func(t int64) string {
		return strconv.FormatInt(t, 10) + "s"
	}

	for _, info := range taskInfos {
		strList := []string{
			info.ID,
			info.Name,
			info.ProcessorID,
			memoryFormat(info.Memory),
			cpuFormat(info.CPU),
			timeFormat(info.CreateAt),
			timeFormat(info.StartAt),
			secondFormat(info.RunTime),
		}
		table.Append(strList)
	}

	return table, nil
}

func getTaskInfo() ([]model.TaskInfo, error) {
	request := gorequest.New()
	_, body, errs := request.Get(config.Config.BaseURL + "/task/info").End()
	if errs != nil {
		return nil, errs[0]
	}
	var infoResp model.TaskInfoResp
	err := json.Unmarshal([]byte(body), &infoResp)
	if err != nil {
		return nil, err
	}
	return infoResp.TaskInfo, nil
}

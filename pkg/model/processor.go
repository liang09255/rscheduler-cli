package model

type ProcessorInfoResp struct {
	BaseResp
	ProcessorInfo []ProcessorInfo `json:"processor_info"`
}

type ProcessorInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PID          int    `json:"pid"`
	TotalTaskNum uint32 `json:"total_task_num"`
	Running      bool   `json:"running"`
	TaskID       string `json:"task_id"`
}

package model

type TaskInfoResp struct {
	BaseResp
	TaskInfo []TaskInfo `json:"task_info"`
}

type TaskInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ProcessorID string `json:"processor_id"`
	Memory      uint64 `json:"memory"`
	CPU         uint64 `json:"cpu"`
	CreateAt    int64  `json:"create_at"`
	StartAt     int64  `json:"start_at"`
	RunTime     int64  `json:"run_time"`
}

package model

type TaskDeleteResp struct {
	BaseResp
	Success bool   `json:"success"`
	Info    string `json:"info"`
}

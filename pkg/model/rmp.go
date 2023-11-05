package model

type ProcessorDeleteResp struct {
	BaseResp
	Success bool   `json:"success"`
	Info    string `json:"info"`
}

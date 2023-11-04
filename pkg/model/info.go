package model

type BaseInfoResp struct {
	BaseResp
	BaseInfo BaseInfo `json:"base_info"`
}

type BaseInfo struct {
	Version string `json:"version"`
}

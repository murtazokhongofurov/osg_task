package models

type FailureInfo struct {
	StatusCode  int         `json:"status_code"`
	Description string      `json:"description"`
	Error       interface{} `json:"error"`
}

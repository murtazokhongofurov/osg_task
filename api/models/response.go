package models

type FailureInfo struct {
	StatusCode  int         `json:"status_code"`
	Description string      `json:"description"`
	Error       interface{} `json:"error"`
}

type Success struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

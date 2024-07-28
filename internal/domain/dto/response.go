package dto

type BaseResponse struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

package models

type OKResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

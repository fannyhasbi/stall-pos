package common

type commonResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

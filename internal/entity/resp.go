package entity

// api response data struct

type DataWrap struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Refer   interface{} `json:"refer,omitempty"`
}

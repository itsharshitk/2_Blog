package model

type Resp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`

	Data any `json:"data,omitempty"`

	ErrorDetails any `json:"error_details,omitempty"`
}

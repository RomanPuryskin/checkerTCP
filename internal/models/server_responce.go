package models

type Responce struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type ResponeError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

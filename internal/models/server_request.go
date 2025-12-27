package models

type Request struct {
	IP      string `json:"ip" validate:"required"`
	Port    string `json:"port" validate:"required"`
	Timeout string `json:"timeout" validate:"required"`
}

package models

type HttpModels struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any  `json:"data"`
}

type DATA struct {
	Data any `json:"data"`
}
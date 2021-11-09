package handler

import "github.com/IamNator/iot-wind/model"

type resp struct {
	Status bool `json:"status"`
	Code   int  `json:"code"`
	Data   data `json:"data"`
}

type data struct {
	Current Values       `json:"current"`
	Log     []*model.Log `json:"log"`
}

type Values struct {
	Speed     int    `json:"speed"`
	Dir       string `json:"dir"`
	CreatedAt string `json:"created_at"`
}

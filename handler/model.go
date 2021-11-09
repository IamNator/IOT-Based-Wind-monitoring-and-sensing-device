package handler

import (
	"time"

	"github.com/IamNator/iot-wind/model"
)

type resp struct {
	Status bool `json:"status"`
	Code   int  `json:"code"`
	Data   data `json:"data"`
}

type data struct {
	Current Log   `json:"current"`
	Log     []Log `json:"log"`
}

var TimeFormat = time.Stamp

func ModelLogsToLogSlice(logs []*model.Log) []Log {
	values := make([]Log, 0)
	for _, l := range logs {
		values = append(values, Log{
			ID:        l.ID,
			Speed:     l.Speed,
			Dir:       l.Dir,
			CreatedAt: l.CreatedAt.Format(TimeFormat),
		})
	} //

	return values
}

type Log struct {
	ID        uint    `form:"id" json:"id"`
	Speed     float32 `form:"speed" json:"speed"`
	Dir       string  `form:"dir" json:"dir"`
	CreatedAt string  `form:"created_at" json:"created_at"`
}

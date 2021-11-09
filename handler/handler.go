package handler

import (
	"time"

	"github.com/IamNator/iot-wind/model"

	"github.com/IamNator/iot-wind/storage"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logStorage storage.LogDatabase
}

type Interface interface {
	Get(ctx *gin.Context)
	POST(ctx *gin.Context)
}

func New(s *storage.Storage) Interface {
	return &Handler{
		logStorage: storage.NewLog(s),
	}
}

func (h *Handler) Get(ctx *gin.Context) {

	current, er := h.logStorage.FindRecent()
	if er != nil {
		ctx.JSONP(422, er.Error())
		return
	}

	logs, err := h.logStorage.FindAllLogs(1, 105)
	if err != nil {
		ctx.JSONP(422, err.Error())
		return
	}

	str := resp{
		Status: true,
		Code:   200,
		Data: data{
			Current: Values{
				Speed:     current.Speed,
				Dir:       current.Dir,
				CreatedAt: current.CreatedAt.Format(time.Stamp),
			},
			Log: logs,
		},
	}

	ctx.JSONP(200, str)
}

func (h *Handler) POST(ctx *gin.Context) {
	var log model.Log
	if er := ctx.BindJSON(&log); er != nil {
		ctx.JSONP(400, er.Error())
		return
	}

	if er := h.logStorage.CreateLog(log); er != nil {
		ctx.JSONP(422, er.Error())
		return
	}

	ctx.JSONP(201, "successfully created")
}

package handler

import (
	"github.com/IamNator/iot-wind/model"
	"github.com/IamNator/iot-wind/pkg/environment"
	"github.com/IamNator/iot-wind/storage"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handler struct {
	logStorage storage.LogDatabase
	env        *environment.Env
}

type Interface interface {
	Get(ctx *gin.Context)
	POST(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

func New(s *storage.Storage, env *environment.Env) Interface {
	return &Handler{
		logStorage: storage.NewLog(s),
		env:        env,
	}
}

func (h *Handler) Get(ctx *gin.Context) {

	recentLog, er := h.logStorage.FindRecent()
	if er != nil {
		ctx.JSONP(422, er.Error())
		return
	}

	logs, err := h.logStorage.FindAllLogs(1, 105)
	if err != nil {
		ctx.JSONP(422, gin.H{"error": err.Error()})
		return
	}

	str := resp{
		Status: true,
		Code:   200,
		Data: data{
			Current: Log{
				ID:        recentLog.ID,
				Speed:     recentLog.Speed,
				Dir:       recentLog.Dir,
				CreatedAt: recentLog.CreatedAt.Format(TimeFormat),
			},
			Log: ModelLogsToLogSlice(logs),
		},
	}

	ctx.JSONP(200, gin.H{"message": str}) //
}

func (h *Handler) POST(ctx *gin.Context) {
	var log Log
	if er := ctx.ShouldBindQuery(&log); er != nil {
		ctx.JSONP(400, gin.H{"error": er.Error()})
		return
	}

	if log.Speed == 0 || log.Dir == "" {
		ctx.JSONP(400, gin.H{"error": "speed and dir are required"})
		return
	}

	if er := h.logStorage.CreateLog(model.Log{
		Speed: log.Speed,
		Dir:   log.Dir, //
	}); er != nil {
		ctx.JSONP(422, gin.H{"error": er.Error()})
		return
	}

	ctx.JSONP(201, gin.H{"message": "successfully created"})
}

func (h *Handler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSONP(400, errors.New("id is required"))
		return
	}

	secret := ctx.Param("secret")
	if secret != h.env.Get("SECRET_KEY") {
		ctx.JSONP(401, gin.H{"error": "unauthorized access"})
	}

	if er := h.logStorage.DeleteByID(id); er != nil {
		ctx.JSONP(422, gin.H{"error": er.Error()})
		return
	}

	ctx.JSONP(201, gin.H{"message": "successfully deleted"})
}

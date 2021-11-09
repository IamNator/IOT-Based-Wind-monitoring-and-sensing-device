package main

import (
	"net/http"
	"os"

	"github.com/IamNator/iot-wind/handler"
	"github.com/IamNator/iot-wind/pkg/environment"
	"github.com/IamNator/iot-wind/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	env, er := environment.New()
	if er != nil {
		panic(er.Error())
	}
	store := storage.New(env)
	handlers := handler.New(store)
	router := gin.Default()

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSONP(200, "no excuses!!")
	})

	router.GET("/get", handlers.Get)
	router.POST("/add", handlers.POST)

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/assets/img/favicon.png")
	router.StaticFile("/", "./static/home.html")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	_ = http.ListenAndServe(":"+port, router)
}

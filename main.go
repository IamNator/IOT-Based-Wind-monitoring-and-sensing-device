package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IamNator/iot-wind/pkg/middleware"

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
	if er := storage.Migration(store); er != nil {
		panic(er.Error())
	} //
	handlers := handler.New(store, env)
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSONP(200, "no excuses!!")
	}) //

	router.GET("/api/get", handlers.Get)
	router.POST("/api/add", handlers.POST)
	router.DELETE("/api/delete", handlers.POST)

	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/assets/img/favicon.png")
	router.StaticFile("/", "./static/home.html")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("serving on port :%s\n", port)
	log.Printf("http://localhost:%s", port)
	_ = http.ListenAndServe(":"+port, router)
}

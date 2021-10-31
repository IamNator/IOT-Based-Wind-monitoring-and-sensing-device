package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main(){

	router := gin.Default()
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/assets/img/favicon.png")
	router.StaticFile("/", "./static/home.html")
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSONP(200, "we are live!!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_ = http.ListenAndServe(":"+port, router)
}
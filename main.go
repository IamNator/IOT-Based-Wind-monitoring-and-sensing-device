package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){

	router := gin.Default()
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/assets/img/favicon.png")
	router.StaticFile("/", "./static/home.html")
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSONP(200, "we are live!!")
	})
	_ = http.ListenAndServe(":3000", router)
}
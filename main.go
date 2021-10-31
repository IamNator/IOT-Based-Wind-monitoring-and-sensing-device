package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"time"
)


type resp struct {
	Status bool `json:"status"`
	Code int `json:"code"`
	Data data `json:"data"`
}

type data struct{
	Current Values `json:"current"`
	Log []Values `json:"log"`
}

type Values struct {
	Speed int `json:"speed"`
	Dir string `json:"dir"`
	CreatedAt string `json:"created_at"`
}

func main(){
	router := gin.Default()
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/assets/img/favicon.png")
	router.StaticFile("/", "./static/home.html")
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSONP(200, "we are live!!")
	})

	router.GET("/get", func(ctx *gin.Context) {
		timeLayout := time.ANSIC
		str := resp{
			Status: true,
			Code:   200,
			Data:   data{
				Current: Values{
					Speed:     (rand.Int()%70),
					Dir:       "North",
					CreatedAt: time.Now().Format(time.ANSIC)},
				Log:     []Values{
					{
						Speed:     (rand.Int()%70),
						Dir:       "North",
						CreatedAt: time.Now().Format(timeLayout)	,
					},
					{
						Speed:     (rand.Int()%70),
						Dir:       "North",
						CreatedAt: time.Now().Add(time.Minute * -5).Format(timeLayout)	,
					},
					{
						Speed:     (rand.Int()%70),
						Dir:       "North",
						CreatedAt: time.Now().Add(time.Minute * -20).Format(timeLayout)	,
					},
				},
			},
		}

		ctx.JSONP(200, str )
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_ = http.ListenAndServe(":"+port, router)
}
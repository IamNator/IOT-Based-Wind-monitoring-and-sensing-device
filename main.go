package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/structpb"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

		str := resp{
			Status: true,
			Code:   200,
			Data:   data{
				Current: Values{
					Speed:     (rand.Int()%70),
					Dir:       "North",
					CreatedAt: time.Now().String(),
				},
				Log:     []Values{
					{
						Speed:     (rand.Int()%70),
						Dir:       "North",
						CreatedAt: time.Now().String(),
					},
					{
						Speed:     (rand.Int()%70),
						Dir:       "North",
						CreatedAt: time.Now().Add(time.Minute * -5).String(),
					},
					{
						Speed:     (rand.Int()%70),
						Dir:       "North",
						CreatedAt: time.Now().Add(time.Minute * -20).String(),
					},
				},
			},
		}
		buf := strings.NewReader(str)
		h := make(map[string]interface{})
		_ = json.NewDecoder(buf).Decode(&h)

		ctx.JSONP(200, h)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_ = http.ListenAndServe(":"+port, router)
}
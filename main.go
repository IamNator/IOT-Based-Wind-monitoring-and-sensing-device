package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func main(){



	router := gin.Default()
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./static/assets/img/favicon.png")
	router.StaticFile("/", "./static/home.html")
	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSONP(200, "we are live!!")
	})

	router.GET("/get", func(ctx *gin.Context) {
		str := `{
  "status": true,
  "code": 200,
  "data": {
      "current":  {
          "speed": 34.89,
          "dir": "56.00 degree north",
          "created_at": "10:00 AM"
      },
      "log": [
      {
          "speed": 34.89,
          "dir": "56.00 degree north",
          "created_at": "10:00 AM"
      },
       {
          "speed": 84.89,
          "dir": "59.00 degree south",
          "created_at": "9:45 AM"
      },
       {
          "speed": 89.89,
          "dir": "79.00 degree north",
          "created_at": "8:00 AM"
      }
  ]
  }
}`
		buf := strings.NewReader(str)
		h := make(map[string]interface{})
		json.NewDecoder(buf).Decode(&h)

		ctx.JSONP(200, h)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_ = http.ListenAndServe(":"+port, router)
}
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	if err := Init(InitGin); err != nil {
		panic(err)
	}
}

func Init(fx ...func() (err error)) (err error) {
	for _, f := range fx {
		if f != nil {
			if err = f(); err != nil {
				return
			}
		}
	}
	return
}

func InitGin() (err error) {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Logger())
	g.Any("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	g.GET("/ws", func(c *gin.Context) {
		UpgradeToWebsocket(Ws1, c.Writer, c.Request)
	})
	err = g.Run("127.0.0.1:9630")
	return
}

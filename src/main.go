package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go StartHttpServer()

	fmt.Println("wenbin test data begin----------")
	fmt.Println(genesis)
	fmt.Println("wenbin test data end  ----------")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("stop signal received.")
}

func StartHttpServer() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/genesis.json", func(c *gin.Context) {
		c.String(http.StatusOK, genesis)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method

		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
	}
}

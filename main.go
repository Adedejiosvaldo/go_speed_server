package main

import "github.com/gin-gonic/gin"

// uploads

// downloads

// pings

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}

package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// uploads

func uploadsTest(ctx *gin.Context) {
	timeStarted := time.Now()

	sampleLoad := make([]byte, 100*1024*1024)

	diffInTime := time.Duration(timeStarted)

}

// downloads

func 

// pings

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()
}

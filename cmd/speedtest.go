package cmd

import (
	"time"

	"github.com/gin-gonic/gin"
)

func downloadSpeedTest(ctx *gin.Context) {

}
func uploadsTest(ctx *gin.Context) {

}
func pingTest(ctx *gin.Context) {
	start := time.Now()

	ctx.JSON(200, gin.H{"message": "pong", "latency": time.Since(start).String()})
}

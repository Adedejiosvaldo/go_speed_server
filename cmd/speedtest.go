package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func downloadSpeedTest(ctx *gin.Context, filePath string) {

	file, err := os.Open(filePath)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	start := time.Now()

	ctx.Stream(func(w io.Writer) bool {
		_, err = io.Copy(w, file)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return false
		}
		return false
	})

	timeElapsed := time.Since(start)

	fileInfo, err := file.Stat()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	downloadSpeed := float64(fileInfo.Size()) / 1024 / 1024 / timeElapsed.Seconds()
	ctx.JSON(200, gin.H{
		"download_time":  timeElapsed.String(),
		"download_speed": fmt.Sprintf("%.2f MB/s", downloadSpeed),
		"file_size":      fmt.Sprintf("%.2f MB/s", float64(fileInfo.Size())),
	})
}
func uploadsTest(ctx *gin.Context) {

}
func pingTest(ctx *gin.Context) {
	start := time.Now()

	ctx.JSON(200, gin.H{"message": "pong", "latency": time.Since(start).String()})
}

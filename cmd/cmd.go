package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func DownloadSpeedTestHandler(filePath string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		DownloadSpeedTest(ctx, filePath)
	}
}

func DownloadSpeedTest(ctx *gin.Context, filePath string) {
	// Open the file to be streamed
	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close() // Ensure the file is closed once the function exits

	// Get file info for size calculation
	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the Content-Type header
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))

	// Start measuring the download time (on the server side)
	start := time.Now()

	// Stream the file to the client
	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		// If any error occurs during streaming, log the error
		fmt.Printf("Error during streaming: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate the download time and speed on the server side (not client-side)
	timeElapsed := time.Since(start)
	downloadSpeed := float64(fileInfo.Size()) / 1024 / 1024 / timeElapsed.Seconds() // Speed in MB/s

	// Send back the download stats (for server side calculation)
	ctx.JSON(http.StatusOK, gin.H{
		"download_time":  timeElapsed.String(),
		"download_speed": fmt.Sprintf("%.2f MB/s", downloadSpeed),
		"file_size":      fmt.Sprintf("%.2f MB", float64(fileInfo.Size())/1024/1024),
	})
}

func UploadsTest(ctx *gin.Context, fileName string) {

	tempFile, err := os.Open(fileName)

	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	defer tempFile.Close()

	start := time.Now()

	destFile, err := os.Create("uploaded_Testfilt.txt")

	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, tempFile)

	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	timeElapsed := time.Since(start)

	fileInfo, err := tempFile.Stat()
	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	uploadSpeed := float64(fileInfo.Size()) / 1024 / 1024 / timeElapsed.Seconds()
	ctx.JSON(200, gin.H{
		"upload_time":  timeElapsed.String(),
		"upload_speed": fmt.Sprintf("%.2f MB/s", uploadSpeed),
		"file_size":    fmt.Sprintf("%.2f MB", float64(fileInfo.Size())/1024/1024),
	})
}
func PingTest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong", "latency": time.Since(start).String()})
}

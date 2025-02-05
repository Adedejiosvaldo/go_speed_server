// package main

// import (
// 	"fmt"

// 	"github.com/adedejiosvaldo/go_speedtest_server/cmd"
// 	"github.com/adedejiosvaldo/go_speedtest_server/helper"
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	err := helper.GenerateLargeFile("testfile.txt", 10)

// 	if err != nil {
// 		fmt.Println("Error generating test file:", err)
// 		return
// 	}

// 	r := gin.Default()

// 	r.GET("/ping", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{"message": "pong"})
// 	})

// 	r.GET("/pings", cmd.PingTest)
// 	r.GET("/download-speed", cmd.DownloadSpeedTestHandler("testfile.txt"))

// 	r.Run()
// }

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.Default())

	// Ping endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Download test endpoint
	router.GET("/download", func(c *gin.Context) {
		// Set default size to 10MB
		size := 10 * 1024 * 1024 // 10 MB
		if customSize := c.Query("size"); customSize != "" {
			// Parse custom size from query parameter
			var err error
			size, err = parseInt(customSize)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid size parameter"})
				return
			}
		}

		// Generate random data
		data := generateRandomData(size)

		// Set headers for octet-stream
		c.DataFromReader(
			http.StatusOK,
			int64(size),
			"application/octet-stream",
			bytes.NewReader(data),
			map[string]string{
				"Content-Disposition": `attachment; filename="test.data"`,
			},
		)
	})

	// Upload test endpoint
	// router.POST("/upload", func(c *gin.Context) {
	// 	// Discard the uploaded data but count the bytes
	// 	size, err := c.Request.Body.(io.ReaderFrom).ReadFrom(io.Discard)
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "upload failed"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"size": size,
	// 	})
	// })

	// Start server
	router.Run(":8080")
}

func generateRandomData(size int) []byte {
	data := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	rand.Read(data)
	return data
}

func parseInt(s string) (int, error) {
	var size int
	_, err := fmt.Sscanf(s, "%d", &size)
	return size, err
}

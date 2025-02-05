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
	"fmt"

	"github.com/adedejiosvaldo/go_speedtest_server/cmd"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.Default())

	router.GET("/ping", cmd.PingTest)
	router.POST("/upload", cmd.UploadsTest())
	router.GET("/download", cmd.DownloadSpeedTestHandler("testfile.dat")) // Change file path

	// Start server
	fmt.Println("Speedtest server running on port 8080")
	router.Run(":8080")

	router.Run(":8080")
}

package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type Pheremone struct {
	Id    string `json:"id" binding:"required"`
	From  string `json:"from" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	go StatsWorker()
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/data", func(c *gin.Context) {
		var json Pheremone

		c.Bind(&json) // This will infer what binder to use depending on the content-type header.

		fmt.Println("Got request :", json)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}

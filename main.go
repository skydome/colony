package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	gin "github.com/gin-gonic/gin"
)

type Pheremone struct {
	At    time.Time `json:"at"`
	Ant   string    `json:"ant" binding:"required"`
	Type  string    `json:"type" binding:"required"`
	Value string    `json:"value" binding:"required"`
}

func main() {
	f, err := os.OpenFile("/tmp/cassa.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	var buf bytes.Buffer
	log.New(&buf, "logger: ", log.Lshortfile)
	log.Println("This is a test log entry")

	ConfigRuntime()
	Initialize()
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
	//	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/telemetry/plant/id/:id", func(c *gin.Context) {
		var pheremone Pheremone
		plantId := c.Params.ByName("id")
		c.Bind(&pheremone)

		pheremone.At = time.Now()
		log.Println("Got request :", pheremone)
		WriteToCassandra(plantId, pheremone)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.GET("/telemetry/plant/id/:id", func(c *gin.Context) {

	})

	r.Run(":8080")
}

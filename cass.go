package main

import (
	"log"

	"github.com/hailocab/gocassa"
)

var keySpace gocassa.KeySpace

var plantMap map[string]gocassa.MultimapTable

func Initialize() {
	log.Println("Initializing GoCassa")

	var err error
	connection, _ := gocassa.Connect([]string{"api.skydome.io"}, "", "")
	connection.CreateKeySpace("test")
	keySpace, err = gocassa.ConnectToKeySpace("test", []string{"api.skydome.io"}, "", "")
	if err != nil {
		panic(err)
	}

	plantMap = map[string]gocassa.MultimapTable{}
}

func WriteToCassandra(plantId string, pheremone Pheremone) {
	log.Println("Writing pheremone: ", pheremone)
	log.Println("With plantId: ", plantId)
	plant := plantMap[plantId]
	if plant == nil {
		plant = keySpace.MultimapTable(plantId, "Ant", "At", Pheremone{})
		log.Print("creating table : ", plantId)

		err := plant.Create()
		if err != nil {
			log.Println("Error:", err)
		}
		plantMap[plantId] = plant
	}
	err := plant.Set(pheremone).Run()
	if err != nil {
		panic(err)
	}
}

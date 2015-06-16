package main

import (
	"log"

	"github.com/hailocab/gocassa"
)

var keySpace gocassa.KeySpace

var plantMap map[string]gocassa.MapTable

func Initialize() {
	log.Println("Initializing GoCassa")

	var err error
	keySpace, err = gocassa.ConnectToKeySpace("test", []string{"api.skydome.io"}, "", "")
	if err != nil {
		panic(err)
	}
	plantMap = map[string]gocassa.MapTable{}
}

func WriteToCassandra(plantId string, pheremone Pheremone) {
	log.Println("Writing pheremone: ", pheremone)
	log.Println("With plantId: ", plantId)
	plant := plantMap[plantId]
	if plant == nil {
		plant = keySpace.MapTable(plantId, "Ant", Pheremone{})
		log.Print("creating table : ", plantId)
		err := plant.Create()
		if err != nil {
			log.Println("Error:", err)
			return
		} else {
			plantMap[plantId] = plant
		}
	}
	err := plant.Set(pheremone).Run()
	if err != nil {
		panic(err)
	}

	//	result := Pheremone{}
	//	if err := salesTable.Read("abcdefg", &result).Run(); err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(result)
}

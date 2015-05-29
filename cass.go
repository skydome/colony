package main

import (
	"fmt"

	"github.com/hailocab/gocassa"
)

var keySpace gocassa.KeySpace

var plantMap map[string]gocassa.MapTable

func Initialize() {
	fmt.Println("Initializing GoCassa")
	var err error
	keySpace, err = gocassa.ConnectToKeySpace("test", []string{"127.0.0.1"}, "", "")
	if err != nil {
		panic(err)
	}
	plantMap = map[string]gocassa.MapTable{}
}

func WriteToCassandra(plantId string, pheremone Pheremone) {
	fmt.Println("Writing pheremone: ", pheremone)
	fmt.Println("With plantId: ", plantId)
	plant := plantMap[plantId]
	if plant == nil {
		plant = keySpace.MapTable(plantId, "Ant", Pheremone{})
		fmt.Print("creating table : ", plantId)
		err := plant.Create()
		if err != nil {
			fmt.Println("Error:", err)
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

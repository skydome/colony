package main

import (
	"fmt"

	"github.com/hailocab/gocassa"
)

var salesTable gocassa.MapTable

func Initialize() {
	fmt.Println("Initializing GoCassa")
	keySpace, err := gocassa.ConnectToKeySpace("test", []string{"api.skydome.io"}, "", "")
	if err != nil {
		panic(err)
	}
	salesTable = keySpace.MapTable("pheremone", "Id", Pheremone{})
	// Create the table - we ignore error intentionally
	salesTable.Create()
}

func WriteToCassandra(pheremone Pheremone) {
	fmt.Println("Writing pheremone: ", pheremone)
	err := salesTable.Set(pheremone).Run()
	if err != nil {
		panic(err)
	}

	result := Pheremone{}
	if err := salesTable.Read("abcdefg", &result).Run(); err != nil {
		panic(err)
	}
	fmt.Println(result)
}

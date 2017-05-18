package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name string
	Age int
	Height int
	Weight int
}

func main() {
	gman := person{
		Name:"Goatman",
		Age: 2000,
		Height: 72,
		Weight: 230,
	}

	marsh, error := json.Marshal(gman)
	if error != nil {
		fmt.Println("error:", error)
	}

	os.Stdout.Write(marsh)
}


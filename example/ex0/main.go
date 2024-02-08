package main

import (
	"fmt"

	"github.com/flazhgrowth/fgmapstructparser/lib"
)

type DestStruct struct {
	ID   int64  `justmap:"id"`
	Name string `justmap:"name"`
}

func main() {
	// source data
	src := map[string]interface{}{
		"id":   123,
		"name": "shn ndr nt",
	}

	dest := &DestStruct{}
	if err := lib.New("justmap").Parse(src, dest); err != nil {
		panic(err)
	}
	fmt.Println("dest.ID: ", dest.ID)
	fmt.Println("dest.Name: ", dest.Name)
}

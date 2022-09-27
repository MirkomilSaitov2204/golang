package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name       string     `json:"name"`
	ScienceId  int        `json:"science_Id"`
	IsWorking  bool       `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	prof1 := Person{
		Name:      "BOB",
		ScienceId: 123123,
		IsWorking: true,
		University: University{
			Name: "TEXNAR",
			City: "TASHKENT",
		},
	}

	byteArr, err := json.MarshalIndent(prof1, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))

	err = os.WriteFile("output.json", byteArr, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

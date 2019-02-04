package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ModelStruct struct {
	Name string `json:"name"` //`json:"_"` BERFUNGSI UNTUK CUSTOM FIELD
	Age  int    `json:"age"`
}

func main() {
	var model ModelStruct
	userSatu := ModelStruct{Name: "JAKA WIDADA", Age: 12}
	userDua := []byte(`{"Name": "PARBAWA SUBIANTA", "Age": 12}`)

	// ENCODE JSON
	jsonSatu, err := json.Marshal(userSatu)
	if err != nil {
		errors.New("ERROR !")
	}
	fmt.Println(jsonSatu)         // NO READABLE
	fmt.Println(string(jsonSatu)) //KONVERSI STRING
	fmt.Println("=====================================================")

	// DECODE JSON
	err = json.Unmarshal(userDua, &model)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(model)
}

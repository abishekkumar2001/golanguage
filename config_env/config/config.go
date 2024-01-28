package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Mobile struct {
	Name string `json:"name"`
	Ram  int    `json:"ram"`
	Port int    `json:"port"`
}

func mobileConfig(filename string) (*Mobile, error) { //The return type is pointer to a structure and an error
	file, err := os.Open(filename) //Returns file pointer and an error
	if err != nil {
		fmt.Printf("Error opening file : ", err)
	}

	defer file.Close()

	config := &Mobile{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil

}

func main() {
	fmt.Println("Config in Go")
	config, err := mobileConfig("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mobile Name is", config.Name)
	fmt.Printf("Ram is %d \n", config.Ram)
	fmt.Println("Port is ", config.Port)
}

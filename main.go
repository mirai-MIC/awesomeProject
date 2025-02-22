package main

import "C"
import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Port int    `json:"port"`
	Env  string `json:"env"`
}

func LoadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func main() {
<<<<<<< HEAD
    getTime()
    new1()
=======
	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	println("Port:", config.Port, "Env:", config.Env)
>>>>>>> f13322963009b7a7f25d44dc3019ece1d1b30650
}
func getTime() {
	t := time.Now()
	fmt.Println(t)
}

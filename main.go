package main

import "C"
import (
	"encoding/json"
	"fmt"
	"log"
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
	config, err := LoadConfig("config.json")
	if err != nil {
		panic(err)
	}
	log.Fatalln("Port:", config.Port, "Env:", config.Env)
}
func getTime() {
	t := time.Now()
	fmt.Println(t)
}

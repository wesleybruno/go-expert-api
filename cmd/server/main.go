package main

import "github.com/wesleybruno/go-expert-api/configs"

func main() {
	configs, err := configs.LoadConfig("./../../")
	if err != nil {
		panic(err)
	}
	print(configs.DBDriver)
}

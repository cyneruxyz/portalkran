package main

import (
	"flag"
	"os"

	"portalkran/internal/app/portalkran"
	"portalkran/internal/pkg/config"
)

func main() {
	configFilePath := flag.String("source", "./portalkran.yaml", "configuration file path (string type)")

	file, err := os.ReadFile(*configFilePath)
	if err != nil {
		panic(err.Error())
	}

	err = portalkran.Run(config.New(file))
	if err != nil {
		panic(err.Error())
	}
}

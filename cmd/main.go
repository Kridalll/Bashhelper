package main

import "github.com/Kridalll/Bashhelper/internal/app"

const configFilePath = "./config/config.yaml"

func main() {
	app.Run(configFilePath)
}

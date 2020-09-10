package main

import (
	"fmt"
	"game/app"
	"game/config"
)

func main() {
	config := config.GetConfig()

	fmt.Println("Server Started | Port", config.Port)
	app := &app.App{}
	app.Initialize(config)
	app.Run(config.Port)

}

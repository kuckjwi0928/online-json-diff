package main

import (
	"fmt"
	"online-json-diff/configs"
	"online-json-diff/internal"
	"online-json-diff/pkg/service"
)

func main() {
	app := internal.NewApp(service.NewServiceContainer())
	app.Run(fmt.Sprintf(":%d", configs.Server.Port))
}

package main

import (
	"api.default.marincor.com/app/appinstance"
	"api.default.marincor.com/pkg/app"
)

func main() {
	app.ApplicationInit()
	defer appinstance.Data.DB.Close()

	appinstance.Data.Server = route()

	// Listening to Server
	app.Setup()
}

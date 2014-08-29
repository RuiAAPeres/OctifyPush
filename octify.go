package main

import (
	"log"

	"github.com/ruiaaperes/octify/controller"
	"github.com/zenazn/goji"
)

const (
	userURLPath = "/v1/user"
)

func main() {

	controller, error := controller.NewController()
	if error != nil {
		log.Fatal(error)
	}

	defer controller.CloseSession()

	goji.Post(userURLPath, controller.RegisterUser)
	goji.Get(userURLPath, controller.RegisteredUser)
	goji.Delete(userURLPath, controller.UnregisterUser)

	goji.Serve()
}

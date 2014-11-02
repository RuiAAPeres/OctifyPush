package main

import (
	"log"

	"github.com/ruiaaperes/octify/controller"
	"github.com/zenazn/goji"
)

const (
	userURLPath     = "/v1/user"
	userPlaceholder = "/:username"
)

func main() {

	controller, error := controller.NewController()
	if error != nil {
		log.Fatal(error)
	}

	controller.StartPush()
	defer controller.Close()

	goji.Post(userURLPath, controller.RegisterUser)
	goji.Get(userURLPath+userPlaceholder, controller.RegisteredUser)
	goji.Delete(userURLPath+userPlaceholder, controller.UnregisterUser)

	goji.Serve()
}

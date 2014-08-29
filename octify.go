package main

import (
	"log"

	"github.com/ruiaaperes/octify/controller"
	"github.com/zenazn/goji"
)

const (
	userURLPath = "/user"
)

func main() {

	controller, error := controller.NewController()
	if error != nil {
		  log.Fatal(error)
	}

    defer controller.CloseSession()

	goji.Post(userURLPath, controller.PostUser)
	goji.Get(userURLPath, controller.GetUser)
	goji.Delete(userURLPath, controller.DeleteUser)

	goji.Serve()
}
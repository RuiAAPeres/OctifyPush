package main

import (
	"fmt"
	"log"

	"github.com/ruiaaperes/octify/controller"
	"github.com/zenazn/goji"
)

const (
	userURLPath = "/user"
)

func main() {

    fmt.Println("Listenasasddasing...")

	controller, error := controller.NewController()
	if error != nil {
		  log.Fatal(error)
	}
    fmt.Println("asd...")

    defer controller.CloseSession()


    fmt.Println("Listening...")
	goji.Post(userURLPath, controller.RegisterUser)
	goji.Get(userURLPath, controller.RegisteredUser)
	goji.Delete(userURLPath, controller.UnregisterUser)

	goji.Serve()
}
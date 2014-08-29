package main

import (
	"log"

	"github.com/ruiaaperes/octify/db"

	"github.com/zenazn/goji"
)

// HTTP Handlers

const (
	userURLPath = "/user"
)

func main() {

	controller, error := db.NewController()
	if error != nil {
		  log.Fatal(error)
	}

    defer controller.CloseSession()

	goji.Post(userURLPath, controller.PostUser)
	goji.Get(userURLPath, controller.GetUser)
	goji.Delete(userURLPath, controller.DeleteUser)

	goji.Serve()
}
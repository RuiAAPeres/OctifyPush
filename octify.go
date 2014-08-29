package main

import (
	"os"
	"net/http"

	"github.com/ruiaaperes/octify/utilities"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

func main() {

	session = setupDB()
	defer session.Close()

	goji.Post("/user", postUser)
	goji.Get("/user", getUser)
	goji.Delete("/user", deleteUser)

	goji.Serve()
}

// DB

func setupDB() *mgo.Session {
	uri := os.Getenv("MONGOHQ_URL")
	if uri == "" {
		panic("no DB connection string provided")
	}

	session, err := mgo.Dial(uri)
	utilities.PanicError(err)
	session.SetMode(mgo.Monotonic, true)

	return session
}

// HTTP Handlers

func postUser(c web.C, w http.ResponseWriter, r *http.Request) {

}

func getUser(c web.C, w http.ResponseWriter, r *http.Request) {

}

func deleteUser(c web.C, w http.ResponseWriter, r *http.Request) {

}


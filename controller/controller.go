package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ruiaaperes/octify/model"

	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoURL        = "MONGOHQ_URL"
	octifyDB        = "octify"
	usersCollection = "users"
	userPlaceholder = "username"
)

type Controller struct {
	*mgo.Session
}

func NewController() (*Controller, error) {

	uri := os.Getenv("MONGOHQ_URL")
	fmt.Errorf(uri)

	if uri == "" {
		return nil, fmt.Errorf("no DB connection string provided")
	}

	session, err := mgo.Dial(uri)
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)

	collection := session.DB(octifyDB).C(usersCollection)

	index := mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}
	err = collection.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}

	return &Controller{
		Session: session,
	}, nil
}

// Start Push

func (controller *Controller) StartPush() {

	ticketChannel := time.NewTicker(time.Second * 10).C

	for {
		select {
		case <-ticketChannel:
			collection := controller.Session.DB(octifyDB).C(usersCollection)
			var allUsers []model.User
			collection.Find(bson.M{}).All(&allUsers)

			for _, element := range allUsers {
				fmt.Printf(element.Username)
				fmt.Printf("\n")
			}

		}
	}
}

// HTTP Handlers

func (controller *Controller) RegisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil || user == (model.User{}) {
		http.Error(w, "Bad content", http.StatusInternalServerError)
	}

	collection := controller.Session.DB(octifyDB).C(usersCollection)
	err = collection.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}
}

func (controller *Controller) RegisteredUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user model.User
	collection := controller.Session.DB(octifyDB).C(usersCollection)
	username := c.URLParams[userPlaceholder]

	err := collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(nil)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(user)
	}
}

func (controller *Controller) UnregisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Done Delete")
}

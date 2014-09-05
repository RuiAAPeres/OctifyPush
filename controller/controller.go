package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

	if err = collection.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}

	return &Controller{
		Session: session,
	}, nil
}

// HTTP Handlers

// POST
func (controller *Controller) RegisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil || user == (model.User{}) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad Content"))
	}

	collection := controller.DB(octifyDB).C(usersCollection)

	if err := collection.Insert(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(user)
	}
}

// GET
func (controller *Controller) RegisteredUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user model.User
	collection := controller.Session.DB(octifyDB).C(usersCollection)
	username := c.URLParams[userPlaceholder]

	if err := collection.Find(bson.M{"username": username}).One(&user); err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(user)
	}
}

// DELETE
func (controller *Controller) UnregisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user model.User
	collection := controller.DB(octifyDB).C(usersCollection)
	username := c.URLParams[userPlaceholder]

	if err := collection.Find(bson.M{"username": username}).One(&user); err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if err := collection.RemoveId(user.Id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

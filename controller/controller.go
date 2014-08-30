package controller

import (
	"encoding/json"
	"io"
	"net/http"
	//	"os"

	"github.com/ruiaaperes/octify/model"

	"github.com/zenazn/goji/web"
	"gopkg.in/mgo.v2"
)

const (
	mongoURL = "MONGOHQ_URL"
)

type Controller struct {
	session *mgo.Session
}

func NewController() (*Controller, error) {

	// uri := os.Getenv("MONGOHQ_URL")
	// fmt.Errorf(uri)

	// if uri == "" {
	// 	return nil, fmt.Errorf("no DB connection string provided")
	// }

	// session, err := mgo.Dial(uri)
	// if err != nil {
	// 	return nil, err
	// }
	// session.SetMode(mgo.Monotonic, true)

	return &Controller{
		session: nil,
	}, nil
}

// Close Session

func (controller *Controller) CloseSession() {
	controller.session.Close()
}

// HTTP Handlers

func (controller *Controller) RegisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil || user == (model.User{}) {
		http.Error(w, "Bad content", http.StatusInternalServerError)
	}

	io.WriteString(w, "Done Post")
}

func (controller *Controller) RegisteredUser(c web.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Done Get")
}

func (controller *Controller) UnregisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Done Delete")
}

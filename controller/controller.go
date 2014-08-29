package controller

import (
	//"os"
    "net/http"
 	"fmt"

    //"github.com/ruiaaperes/octify/model"

	"gopkg.in/mgo.v2"
    "github.com/zenazn/goji/web"
)

const (
	mongoURL = "MONGOHQ_URL"
)

type Controller struct {
    session *mgo.Session
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
    fmt.Println(w,"Done Post")
}

func (controller *Controller) RegisteredUser(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Println(w,"Done Get")
}

func (controller *Controller) UnregisterUser(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Println(w,"Done Delete")
}
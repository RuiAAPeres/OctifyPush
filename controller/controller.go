package controller

import (
	"os"
    "net/http"
 	"fmt"

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

    uri := os.Getenv(mongoURL)
    if uri == "" {
        return nil, fmt.Errorf("no DB connection string provided")
    }

    session, err := mgo.Dial(uri)
    if err != nil {
        return nil, err
    }
    session.SetMode(mgo.Monotonic, true)

    return &Controller{
        session: session,
    }, nil
}

// Close Session

func (controller *Controller) CloseSession() {
    controller.session.Close()
}

// HTTP Handlers

func (controller *Controller) PostUser(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) GetUser(c web.C, w http.ResponseWriter, r *http.Request) {

}

func (controller *Controller) DeleteUser(c web.C, w http.ResponseWriter, r *http.Request) {

}
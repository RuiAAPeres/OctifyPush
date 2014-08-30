package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Token    string        `json:"token" bson:"token"`
	Oauth    string        `json:"oauth" bson:"oauth"`
	Username string        `json:"username" bson:"username"`
}

package model

type User struct {
	token    string `json:"token"`
	oauth    string `json:"oauth"`
	username string `json:"username"`
}

package model

type User struct {
	Token    string `json:"token"`
	Oauth    string `json:"oauth"`
	Username string `json:"username"`
}

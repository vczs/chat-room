package model

type User struct {
	Admin int `json:"admin"`
	Password string `json:"password"`
	AdminName string `json:"adminName"`
	State int `json:"state"`
}
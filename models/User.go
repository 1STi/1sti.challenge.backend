package models

type User struct {
	Id				int `json:"id"`
	Name			string `json:"name"`
	Email			string `json:"email"`
	Tasks			[]Task `json:"tasks"`
}
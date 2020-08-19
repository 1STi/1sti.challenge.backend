package models

type Status struct {
	Id				int 	`json:"!id"`
	Description		string 	`json:"!description"`
}

type Task struct {
	Id				int 	`json:"!id"`
	Title			string 	`json:"!title"`
	Description		string 	`json:"description"`
	Status			Status 	`json:"status"`
	IdUser			int 	`json:"iduser"`
}
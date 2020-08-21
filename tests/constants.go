package tests

import (
	"github.com/YtaloWill/1sti.challenge.backend/models"
)

var User = models.User{
	Name: "User",
	Email: "email@test.com",
	Tasks: []models.Task{},
}

var EmptyUser = models.User{
	Name: "",
	Email: "",
	Tasks: []models.Task{},
}

var UserList = []models.User{
	User,
}

var Task = models.Task{
	Title: "Task Test",
	Description: "Task Test Description",
	Status: models.Status{Id: 1},
	IdUser: 1,
}
package controllers_test

import (
	"testing"
	"github.com/YtaloWill/1sti.challenge.backend/controllers"
	"github.com/YtaloWill/1sti.challenge.backend/tests"
)

func TestAddUser(t *testing.T) {
	got := controllers.AddUser(tests.User)
	want := 2
	
	if(got != want){ t.Errorf("Want: %d, got: %d", want, got) }
	
}

func TestGetWrongUserByName(t *testing.T) {
	got := controllers.GetUserByName(tests.EmptyUser.Name)
	want := tests.EmptyUser
	
	if(got.Name != want.Name && got.Email != want.Email){ t.Errorf("Want: %v, got: %v", want, got) }
	if(len(got.Tasks) > 0){ t.Errorf("Excess of tasks. Want: %d tasks, got: %d tasks", 0, len(got.Tasks)) }
}

func TestGetUserByName(t *testing.T) {
	got := controllers.GetUserByName(tests.User.Name)
	want := tests.User
	
	if(got.Name != want.Name && got.Email != want.Email){ t.Errorf("Want: %v, got: %v", want, got) }
	for i, value := range got.Tasks {
		if(value != want.Tasks[i]){ 
			t.Errorf("Wrong tasks from this user: Want: %v, got: %v", want.Tasks[i], value) 
		}
	}
}

func TestGetAllUsers(t *testing.T) {
	got := controllers.GetAllUsers()
	want := 2

	if (len(got) != want) { 
		t.Errorf("Wrong information from this user: Want: %v, got: %v", want, len(got)) 
	}
}

func TestUpdateWrongUserById(t *testing.T) {
	got := controllers.UpdateUserById(999, tests.User)
	want := 0
	
	if(got != want){ t.Errorf("Want: %d, got: %d", want, got) }
}


func TestUpdateUserById(t *testing.T) {
	got := controllers.UpdateUserById(1, tests.User)
	want := 1
	
	if(got != want){ t.Errorf("Want: %d, got: %d", want, got) }
}

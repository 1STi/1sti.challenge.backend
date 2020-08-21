package controllers_test

import (
	"testing"
	"github.com/YtaloWill/1sti.challenge.backend/controllers"
	"github.com/YtaloWill/1sti.challenge.backend/tests"
	"github.com/YtaloWill/1sti.challenge.backend/database"
)

func init(){
	tests.BuildTestDb()
}

func TestAddTask(t *testing.T) {
	_ = database.Db.QueryRow(`INSERT INTO tbusers values(DEFAULT, 'test', 'email@test.com');`)

	got := controllers.AddTask(tests.Task)
	want := 1
	
	if(got != want){ t.Errorf("Want: %d, got: %d", want, got) }
}

func TestUpdateNoExistentTask(t *testing.T) {
	got := controllers.UpdateTaskById(999, tests.Task)
	want := 0

	if(got != want){ t.Errorf("Want: %d, got: %d", want, got) }
}

func TestUpdateTask(t *testing.T) {
	got := controllers.UpdateTaskById(1, tests.Task)
	want := 1

	if(got != want){ t.Errorf("Want: %d, got: %d", want, got) }

	tests.ClearTables()
}
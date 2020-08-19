package controllers

import (
	"log"
	"github.com/YtaloWill/1sti.challenge.backend/models"
	"github.com/YtaloWill/1sti.challenge.backend/database"
)

func getTasksFromUser(idUser int) ([]models.Task) {
	rowsTasks, err := database.Db.Query(`select tbtasks.id, tbtasks.title, tbtasks.description, tbtasks.idstatus, tbstatus.description from tbtasks
	join tbstatus on (tbtasks.idstatus = tbstatus.id)
	where iduser=$1;`, idUser)
	if err != nil { log.Fatal(err) }

	var tasks []models.Task
	var task models.Task
	var status models.Status
	for rowsTasks.Next() {
		rowsTasks.Scan(&task.Id, &task.Title, &task.Description, &status.Id, &status.Description)
		task.Status = status
		tasks = append(tasks, task)
	}
	return tasks
}

func AddTask(newTask models.Task) (ok bool, err error){
	err = database.Db.QueryRow(`INSERT INTO tbtasks 
	VALUES(DEFAULT, $1, $2, $3, $4)
	RETURNING 1;`, newTask.Title, newTask.Description, newTask.IdUser, newTask.Status.Id).Scan(&ok)
	return ok, err
}

func UpdateTaskById(oldTaskId int, newTask models.Task) (ok bool, err error){
	err = database.Db.QueryRow(`UPDATE tbtasks 
	SET title=$1, description=$2, idstatus=$3, iduser=$4
	WHERE id=$5
	RETURNING 1;`, newTask.Title, newTask.Description, newTask.Status.Id, newTask.IdUser, oldTaskId).Scan(&ok)
	return ok, err
}

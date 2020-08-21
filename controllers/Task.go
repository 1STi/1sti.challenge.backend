package controllers

import (
	"github.com/YtaloWill/1sti.challenge.backend/models"
	"github.com/YtaloWill/1sti.challenge.backend/database"
)

func getTasksFromUser(idUser int) ([]models.Task) {
	rowsTasks, _ := database.Db.Query(`select tbtasks.id, tbtasks.title, tbtasks.description, tbtasks.idstatus, tbstatus.description from tbtasks
	join tbstatus on (tbtasks.idstatus = tbstatus.id)
	where iduser=$1;`, idUser)

	var tasks []models.Task
	var task models.Task
	for rowsTasks.Next() {
		rowsTasks.Scan(&task.Id, &task.Title, &task.Description, &task.Status.Id, &task.Status.Description)
		tasks = append(tasks, task)
	}
	return tasks
}

func AddTask(newTask models.Task) (id int){
	_ = database.Db.QueryRow(`INSERT INTO tbtasks 
	VALUES(DEFAULT, $1, $2, $3, $4)
	RETURNING id;`, newTask.Title, newTask.Description, newTask.IdUser, newTask.Status.Id).Scan(&id)
	return id
}

func UpdateTaskById(oldTaskId int, newTask models.Task) (id int){
	_ = database.Db.QueryRow(`UPDATE tbtasks 
	SET title=$1, description=$2, iduser=$3, idstatus=$4
	WHERE id=$5
	RETURNING id;`, newTask.Title, newTask.Description, newTask.IdUser, newTask.Status.Id, oldTaskId).Scan(&id)

	return id
}

package controllers

import (
	"log"
	"github.com/YtaloWill/1sti.challenge.backend/models"
	"github.com/YtaloWill/1sti.challenge.backend/database"
)

func GetUserByName(userName string) (user models.User, err error){
	err = database.Db.QueryRow(`select id, name, email from tbusers where name=$1;`, userName).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil { log.Fatal(err) }

	user.Tasks = getTasksFromUser(user.Id)

	return user, err
}

func GetAllUsers() (users []models.User, err error) {
	rowsUsers, err := database.Db.Query("SELECT id, name, email FROM tbusers;")
	if err != nil { log.Fatal(err) }

	for rowsUsers.Next() {
		var user models.User
		rowsUsers.Scan(&user.Id, &user.Name, &user.Email)
		user.Tasks = getTasksFromUser(user.Id)

		users = append(users, user)
	}

	return users, err
}

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

func AddUser(newUser models.User) (ok bool, err error){
	err = database.Db.QueryRow(`INSERT INTO tbusers 
	VALUES(DEFAULT, $1, $2)
	RETURNING 1;`, newUser.Name, newUser.Email).Scan(&ok)
	return ok, err
}

func UpdateUserById(oldUserId int, newUser models.User) (ok bool, err error){
	err = database.Db.QueryRow(`UPDATE tbusers 
	SET name=$1, email=$2
	WHERE id=$3
	RETURNING 1;`, newUser.Name, newUser.Email, oldUserId).Scan(&ok)
	return ok, err
}

func DeleteUserById(deleteUserId int) (ok bool, err error){
	err = database.Db.QueryRow(`DELETE FROM tbusers WHERE id=$1
	RETURNING 1;`, deleteUserId).Scan(&ok)
	return ok, err
}
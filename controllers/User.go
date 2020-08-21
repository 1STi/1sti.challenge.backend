package controllers

import (
	"github.com/YtaloWill/1sti.challenge.backend/models"
	"github.com/YtaloWill/1sti.challenge.backend/database"
)

func GetUserByName(userName string) (user models.User){
	_ = database.Db.QueryRow(`select id, name, email from tbusers where name=$1;`, userName).Scan(&user.Id, &user.Name, &user.Email)

	user.Tasks = getTasksFromUser(user.Id)

	return user
}

func GetAllUsers() (users []models.User) {
	rowsUsers, _ := database.Db.Query("SELECT id, name, email FROM tbusers;")

	for rowsUsers.Next() {
		var user models.User
		rowsUsers.Scan(&user.Id, &user.Name, &user.Email)
		user.Tasks = getTasksFromUser(user.Id)

		users = append(users, user)
	}

	return users
}

func AddUser(newUser models.User) (id int){
	_ = database.Db.QueryRow(`INSERT INTO tbusers 
	VALUES(DEFAULT, $1, $2)
	RETURNING id;`, newUser.Name, newUser.Email).Scan(&id)
	return id
}

func UpdateUserById(oldUserId int, newUser models.User) (id int){
	_ = database.Db.QueryRow(`UPDATE tbusers 
	SET name=$1, email=$2
	WHERE id=$3
	RETURNING id;`, newUser.Name, newUser.Email, oldUserId).Scan(&id)
	return id
}

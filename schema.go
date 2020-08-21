package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/YtaloWill/1sti.challenge.backend/models"
	"github.com/YtaloWill/1sti.challenge.backend/gqlmodels"
	"github.com/YtaloWill/1sti.challenge.backend/controllers"	
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Query",
	Description: "Query any user by name",
	Fields: graphql.Fields{
		"users": {
			Type: graphql.NewList(gqlmodels.User),
			Description: "Get information about all users and their tasks",
			Args: graphql.FieldConfigArgument{},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return controllers.GetAllUsers(), nil
			},
		},
		"user": {
			Type: gqlmodels.User,
			Description: "Get information about a specifically user and tasks using his name",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "Name of the user to search",
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Args["name"] != nil {
					name := p.Args["name"].(string)
					return controllers.GetUserByName(name), nil
				}

				return nil, fmt.Errorf("you need to specify \"name\" of the user")
			},
		},
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Mutation",
	Description: "Manipulate all information about users and receive a boolean confirmation if mutation done",
	Fields: graphql.Fields{
		"addUser": {
			Type: graphql.Int,
			Description: "Add a new user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "Name of the user",
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Description: "Email of the user",
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {				
				if p.Args["name"] != nil && p.Args["email"] != nil {
					var newUser models.User
					newUser.Name = p.Args["name"].(string)
					newUser.Email = p.Args["email"].(string)
					return controllers.AddUser(newUser), nil
				}
				return false, fmt.Errorf("you need to specify all the information about the new user")
			},
		},
		"updateUserById": {
			Type: graphql.Int,
			Description: "Update information about a user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Id of the user to update information",
					Type: graphql.Int,
				},
				"name": &graphql.ArgumentConfig{
					Description: "User name updated",
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Description: "User email updated",
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Args["id"] != nil && 
				p.Args["name"] != nil &&
				p.Args["email"] != nil {
					id := p.Args["id"].(int)
					var updatedUser models.User
					updatedUser.Name = p.Args["name"].(string)
					updatedUser.Email = p.Args["email"].(string)
					return controllers.UpdateUserById(id, updatedUser), nil
				}
				return nil, fmt.Errorf("you need to specify the \"id\" of the user AND the new user information")
			},
		},
		"deleteUserById": {
			Type: graphql.Boolean,
			Description: "Delete all information about a user TO ETERNITY",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Id of the user to delete information",
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Args["id"] != nil {
					id := p.Args["id"].(int)
					return controllers.DeleteUserById(id)
				}
				return nil, fmt.Errorf("you need to specify the \"id\" of the user")
			},
		},
		"addTask": {
			Type: graphql.Boolean,
			Description: "Add a new task",
			Args: graphql.FieldConfigArgument{
				"title": &graphql.ArgumentConfig{
					Description: "Short description about task",
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Description: "More detailed description about task",
					Type: graphql.String,
				},
				"idUser": &graphql.ArgumentConfig{
					Description: "Id of the user who will execute this task",
					Type: graphql.Int,
				},
				"idStatus": &graphql.ArgumentConfig{
					Description: "Id of the current status from task, 1 for \"a fazer\", 2 for \"fazendo\", 3 for \"feito\"",
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {				
				if p.Args["title"] != nil && 
				p.Args["description"] != nil &&
				p.Args["idUser"] != nil &&
				p.Args["idStatus"] != nil {
					var newTask models.Task
					newTask.Title = p.Args["title"].(string)
					newTask.Description = p.Args["description"].(string)
					newTask.IdUser = p.Args["idUser"].(int)
					newTask.Status.Id = p.Args["idStatus"].(int)
					return controllers.AddTask(newTask), nil
				}
				return false, fmt.Errorf("you need to specify all the information about the new task")
			},
		},
		"updateTaskById": {
			Type: graphql.Int,
			Description: "Update information about a task",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Id of the task to update information",
					Type: graphql.Int,
				},
				"title": &graphql.ArgumentConfig{
					Description: "task title updated",
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Description: "task description updated",
					Type: graphql.String,
				},
				"idUser": &graphql.ArgumentConfig{
					Description: "id of the user who will do the task updated",
					Type: graphql.Int,
				},
				"idStatus": &graphql.ArgumentConfig{
					Description: "id of the current status from task updated",
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Args["id"] != nil &&
				p.Args["title"] != nil && 
				p.Args["description"] != nil &&
				p.Args["idUser"] != nil &&
				p.Args["idStatus"] != nil {
					oldTaskId := p.Args["id"].(int)
					var newTask models.Task
					newTask.Title = p.Args["title"].(string)
					newTask.Description = p.Args["description"].(string)
					newTask.IdUser = p.Args["idUser"].(int)
					newTask.Status.Id = p.Args["idStatus"].(int)
					return controllers.UpdateTaskById(oldTaskId, newTask), nil
				}
				return nil, fmt.Errorf("you need to specify the \"id\" of the task AND the new task information")
			},
		},

	},
})

func GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryType,
		Mutation: MutationType,
	})
}

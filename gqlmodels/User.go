package gqlmodels

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "Represents a task",
	Fields: graphql.Fields{
		"id": {
			Type: graphql.Int,
			Description: "Default id",
		},
		"name": {
			Type: graphql.String,
			Description: "The name of this user",
		},
		"email": {
			Type: graphql.String,
			Description: "The email of this user",
		},
		"tasks": {
			Type: graphql.NewList(Task),
			Description: "List of tasks of this user",
		},
	},
})

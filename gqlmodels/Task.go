package gqlmodels

import (
	"github.com/graphql-go/graphql"
)

var Task = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Task",
	Description: "Represents a task",
	Fields: graphql.Fields{
		"id": {
			Type: graphql.Int,
			Description: "Default id",
		},
		"title": {
			Type: graphql.String,
			Description: "Title of this task",
		},
		"description": {
			Type: graphql.String,
			Description: "More detailed description of this task",
		},
		"status": {
			Type: Status,
			Description: "Current state of this task",
		},
	},
})

var Status = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Status",
	Description: "Represents status of a task",
	Fields: graphql.Fields{
		"id": {
			Type: graphql.Int,
			Description: "Default id",
		},
		"description": {
			Type: graphql.String,
			Description: "Description of the status",
		},
	},
})
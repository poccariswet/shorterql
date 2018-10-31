package handler

import "github.com/graphql-go/graphql"

var shortURLType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ShortURL",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"longURL": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"shortURL": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"count": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"createdAt": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

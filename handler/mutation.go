package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/poccariswet/shorterql/storage"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        shortURLType,
			Description: "Create new short url product",
			Args: graphql.FieldConfigArgument{
				"longURL": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return storage.SaveURL(params.Args["longURL"].(string))
			},
		},
	},
})

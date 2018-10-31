package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/poccariswet/shorterql/storage"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"urlsho": &graphql.Field{
			Type:        shortURLType,
			Description: "fetch the url info by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolveUrlsho,
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(shortURLType),
			Description: "fetch the url info list",
			Resolve:     resolveList,
		},
	},
})

func resolveUrlsho(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(string)
	return storage.FetchURLInfoByID(id)
}

func resolveList(p graphql.ResolveParams) (interface{}, error) {
	return storage.FetchURLInfoList()
}

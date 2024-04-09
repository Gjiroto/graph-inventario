package schema

import (
	"graphql-inventario/mutations"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"refrigerador": &graphql.Field{
			Type:        resolvers.RefrigeradorType,
			Description: "Get Refrigerador by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Implementar la lógica para obtener un refrigerador por su ID
				return nil, nil
			},
		},
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createRefrigerador": mutations.CreateRefrigeradorMutation,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

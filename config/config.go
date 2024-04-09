package config

import (
	"graphql-inventario/schema"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    schema.RootQuery,
	Mutation: schema.RootMutation,
})

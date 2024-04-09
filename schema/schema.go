package schema

import (
	"graphql-inventario/config"
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
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Obtener el ID del refrigerador a consultar
				id, _ := params.Args["id"].(string)

				// Consultar el refrigerador por su ID en la base de datos
				refrigerador := &resolvers.Refrigerador{}
				err := config.DB.QueryRow("SELECT * FROM refrigeradores WHERE id = ?", id).
					Scan(&refrigerador.ID, &refrigerador.Name, &refrigerador.Brand, &refrigerador.Color, &refrigerador.Category)
				if err != nil {
					return nil, err
				}

				return refrigerador, nil
			},
		},
		"allRefrigeradores": mutations.AllRefrigeradoresQuery,
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createRefrigerador": mutations.CreateRefrigeradorMutation,
		"editRefrigerador":   mutations.EditRefrigeradorMutation,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation, // Aquí usas el objeto de mutación combinado
})

package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var AllRefrigeradoresQuery = &graphql.Field{
	Type:        graphql.NewList(resolvers.RefrigeradorType),
	Description: "Get all Refrigeradores",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Consultar todos los refrigeradores de la base de datos
		rows, err := config.DB.Query("SELECT * FROM refrigeradores")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Iterar sobre los resultados y crear una lista de refrigeradores
		refrigeradores := []*resolvers.Refrigerador{}
		for rows.Next() {
			refrigerador := &resolvers.Refrigerador{}
			err := rows.Scan(&refrigerador.ID, &refrigerador.Name, &refrigerador.Brand, &refrigerador.Color, &refrigerador.Category)
			if err != nil {
				return nil, err
			}
			refrigeradores = append(refrigeradores, refrigerador)
		}

		return refrigeradores, nil
	},
}

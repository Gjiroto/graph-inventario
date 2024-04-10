package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var AllTVsQuery = &graphql.Field{
	Type:        graphql.NewList(resolvers.TVType),
	Description: "Get all TVs",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Consultar todos los televisores de la base de datos
		rows, err := config.DB.Query("SELECT * FROM televisores")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Iterar sobre los resultados y crear una lista de televisores
		tvs := []*resolvers.TV{}
		for rows.Next() {
			tv := &resolvers.TV{}
			err := rows.Scan(&tv.ID, &tv.Name, &tv.Brand, &tv.Size)
			if err != nil {
				return nil, err
			}
			tvs = append(tvs, tv)
		}

		return tvs, nil
	},
}

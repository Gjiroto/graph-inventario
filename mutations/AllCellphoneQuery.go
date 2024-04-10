package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var AllCellphonesQuery = &graphql.Field{
	Type:        graphql.NewList(resolvers.CellphoneType),
	Description: "Get all Cellphones",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Consultar todos los teléfonos celulares de la base de datos
		rows, err := config.DB.Query("SELECT * FROM celulares")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Iterar sobre los resultados y crear una lista de teléfonos celulares
		cellphones := []*resolvers.Cellphone{}
		for rows.Next() {
			cellphone := &resolvers.Cellphone{}
			err := rows.Scan(&cellphone.ID, &cellphone.Name, &cellphone.Brand, &cellphone.Model)
			if err != nil {
				return nil, err
			}
			cellphones = append(cellphones, cellphone)
		}

		return cellphones, nil
	},
}

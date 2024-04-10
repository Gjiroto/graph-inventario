package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var EditCellphoneMutation = &graphql.Field{
	Type:        resolvers.CellphoneType,
	Description: "Edit an existing Cellphone",
	Args: graphql.FieldConfigArgument{
		"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"name":  &graphql.ArgumentConfig{Type: graphql.String},
		"brand": &graphql.ArgumentConfig{Type: graphql.String},
		"model": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener los valores de los argumentos
		id, _ := params.Args["id"].(string)
		name, _ := params.Args["name"].(string)
		brand, _ := params.Args["brand"].(string)
		model, _ := params.Args["model"].(string)

		// Actualizar el teléfono celular en la base de datos
		_, err := config.DB.Exec("UPDATE celulares SET name = ?, brand = ?, model = ? WHERE id = ?",
			name, brand, model, id)
		if err != nil {
			return nil, err
		}

		// Obtener el teléfono celular actualizado
		cellphone := &resolvers.Cellphone{
			ID:    id,
			Name:  name,
			Brand: brand,
			Model: model,
		}

		return cellphone, nil
	},
}

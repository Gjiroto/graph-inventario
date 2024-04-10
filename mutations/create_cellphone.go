package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var CreateCellphoneMutation = &graphql.Field{
	Type:        resolvers.CellphoneType,
	Description: "Create a new Cellphone",
	Args: graphql.FieldConfigArgument{
		"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"brand": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"model": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener los valores de los argumentos
		name, _ := params.Args["name"].(string)
		brand, _ := params.Args["brand"].(string)
		model, _ := params.Args["model"].(string)

		// Crear un nuevo teléfono celular
		cellphone := &resolvers.Cellphone{
			ID:    uuid.New().String(),
			Name:  name,
			Brand: brand,
			Model: model,
		}

		// Guardar el teléfono celular en la base de datos
		_, err := config.DB.Exec("INSERT INTO celulares (id, name, brand, model) VALUES (?, ?, ?, ?)",
			cellphone.ID, cellphone.Name, cellphone.Brand, cellphone.Model)
		if err != nil {
			return nil, err
		}

		return cellphone, nil
	},
}

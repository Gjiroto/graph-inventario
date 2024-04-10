package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var EditRefrigeradorMutation = &graphql.Field{
	Type:        resolvers.RefrigeradorType,
	Description: "Edit an existing Refrigerador",
	Args: graphql.FieldConfigArgument{
		"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"name":  &graphql.ArgumentConfig{Type: graphql.String},
		"brand": &graphql.ArgumentConfig{Type: graphql.String},
		"color": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener el ID del refrigerador a editar
		id, _ := params.Args["id"].(string)

		// Obtener los nuevos valores del refrigerador (si se proporcionaron)
		name, _ := params.Args["name"].(string)
		brand, _ := params.Args["brand"].(string)
		color, _ := params.Args["color"].(string)

		// Obtener el refrigerador de la base de datos
		refrigerador := &resolvers.Refrigerador{}
		err := config.DB.QueryRow("SELECT * FROM refrigeradores WHERE id = ?", id).
			Scan(&refrigerador.ID, &refrigerador.Name, &refrigerador.Brand, &refrigerador.Color, &refrigerador.Category)
		if err != nil {
			return nil, err
		}

		// Actualizar los campos si se proporcionaron nuevos valores
		if name != "" {
			refrigerador.Name = name
		}
		if brand != "" {
			refrigerador.Brand = brand
		}
		if color != "" {
			refrigerador.Color = color
		}

		// Actualizar el refrigerador en la base de datos
		_, err = config.DB.Exec("UPDATE refrigeradores SET name = ?, brand = ?, color = ? WHERE id = ?",
			refrigerador.Name, refrigerador.Brand, refrigerador.Color, refrigerador.ID)
		if err != nil {
			return nil, err
		}

		return refrigerador, nil
	},
}

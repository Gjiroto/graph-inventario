package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/graphql-go/graphql"
)

var EditTVMutation = &graphql.Field{
	Type:        resolvers.TVType,
	Description: "Edit an existing TV",
	Args: graphql.FieldConfigArgument{
		"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"name":  &graphql.ArgumentConfig{Type: graphql.String},
		"brand": &graphql.ArgumentConfig{Type: graphql.String},
		"size":  &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener los valores de los argumentos
		id, _ := params.Args["id"].(string)
		name, _ := params.Args["name"].(string)
		brand, _ := params.Args["brand"].(string)
		size, _ := params.Args["size"].(int)

		// Obtener el televisor de la base de datos
		tv := &resolvers.TV{}
		err := config.DB.QueryRow("SELECT * FROM televisores WHERE id = ?", id).
			Scan(&tv.ID, &tv.Name, &tv.Brand, &tv.Size)
		if err != nil {
			return nil, err
		}

		// Actualizar los campos si se proporcionaron nuevos valores
		if name != "" {
			tv.Name = name
		}
		if brand != "" {
			tv.Brand = brand
		}
		if size != 0 {
			tv.Size = size
		}

		// Actualizar el televisor en la base de datos
		_, err = config.DB.Exec("UPDATE televisores SET name = ?, brand = ?, size = ? WHERE id = ?",
			tv.Name, tv.Brand, tv.Size, tv.ID)
		if err != nil {
			return nil, err
		}

		return tv, nil
	},
}

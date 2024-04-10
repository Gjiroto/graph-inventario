package mutations

import (
	"graphql-inventario/config"

	"github.com/graphql-go/graphql"
)

var DeleteTVMutation = &graphql.Field{
	Type:        graphql.String,
	Description: "Delete an existing TV",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener el ID del televisor a eliminar
		id, _ := params.Args["id"].(string)

		// Eliminar el televisor de la base de datos
		_, err := config.DB.Exec("DELETE FROM televisores WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		return "TV eliminado exitosamente", nil
	},
}

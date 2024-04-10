package mutations

import (
	"graphql-inventario/config"

	"github.com/graphql-go/graphql"
)

var DeleteRefrigeradorMutation = &graphql.Field{
	Type:        graphql.String,
	Description: "Delete an existing Refrigerador",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener el ID del refrigerador a eliminar
		id, _ := params.Args["id"].(string)

		// Eliminar el refrigerador de la base de datos
		_, err := config.DB.Exec("DELETE FROM refrigeradores WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		return "Refrigerador eliminado exitosamente", nil
	},
}

package mutations

import (
	"graphql-inventario/config"

	"github.com/graphql-go/graphql"
)

var DeleteCellphoneMutation = &graphql.Field{
	Type:        graphql.String,
	Description: "Delete an existing Cellphone",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener el ID del teléfono celular a eliminar
		id, _ := params.Args["id"].(string)

		// Eliminar el teléfono celular de la base de datos
		_, err := config.DB.Exec("DELETE FROM celulares WHERE id = ?", id)
		if err != nil {
			return nil, err
		}

		return "Cellphone deleted successfully", nil
	},
}

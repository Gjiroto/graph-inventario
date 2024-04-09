package mutations

import (
	"graphql-inventario/resolvers"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var CreateRefrigeradorMutation = &graphql.Field{
	Type:        resolvers.RefrigeradorType,
	Description: "Create a new Refrigerador",
	Args: graphql.FieldConfigArgument{
		"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"brand":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"color":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"category": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// Obtener los valores de los argumentos
		name, _ := params.Args["name"].(string)
		brand, _ := params.Args["brand"].(string)
		color, _ := params.Args["color"].(string)
		category, _ := params.Args["category"].(string)

		// Crear un nuevo refrigerador
		refrigerador := &resolvers.Refrigerador{

			ID:       uuid.New().String(), // Puedes generar un ID único aquí
			Name:     name,
			Brand:    brand,
			Color:    color,
			Category: resolvers.Category(category),
		}

		// Aquí puedes guardar el refrigerador en tu base de datos u otro almacenamiento
		// Por simplicidad, aquí simplemente devolvemos el refrigerador creado
		return refrigerador, nil
	},
}

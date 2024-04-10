package mutations

import (
	"graphql-inventario/config"
	"graphql-inventario/resolvers"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var CreateTVMutation = &graphql.Field{
	Type:        resolvers.TVType,
	Description: "Create a new TV",
	Args: graphql.FieldConfigArgument{
		"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"brand": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"size":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["name"].(string)
		brand, _ := params.Args["brand"].(string)
		size, _ := params.Args["size"].(int)

		tv := &resolvers.TV{
			ID:    uuid.New().String(),
			Name:  name,
			Brand: brand,
			Size:  size,
		}

		_, err := config.DB.Exec("INSERT INTO televisores (id, name, brand, size) VALUES (?, ?, ?, ?)",
			tv.ID, tv.Name, tv.Brand, tv.Size)
		if err != nil {
			return nil, err
		}

		return tv, nil
	},
}

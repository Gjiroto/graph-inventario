package graph

import "graphql-inventario/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	refrigerator []*model.Refrigerator
	tv           []*model.Tv
	cellphone    []*model.Cellphone
}

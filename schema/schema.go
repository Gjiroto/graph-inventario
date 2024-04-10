package schema

import (
	"graphql-inventario/mutations"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"allRefrigeradores": mutations.AllRefrigeradoresQuery,
		"allTvs":            mutations.AllTVsQuery,
		"allCellphones":     mutations.AllCellphonesQuery,
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createRefrigerador": mutations.CreateRefrigeradorMutation,
		"createTv":           mutations.CreateTVMutation,
		"createCellphone":    mutations.CreateCellphoneMutation,
		"editRefrigerador":   mutations.EditRefrigeradorMutation,
		"editTv":             mutations.EditTVMutation,
		"editCellphone":      mutations.EditCellphoneMutation,
		"deleteRefrigerador": mutations.DeleteRefrigeradorMutation,
		"deleteTv":           mutations.DeleteTVMutation,
		"deleteCellphone":    mutations.DeleteCellphoneMutation,
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

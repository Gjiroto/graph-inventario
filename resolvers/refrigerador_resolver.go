package resolvers

import "github.com/graphql-go/graphql"

type Refrigerador struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Brand    string   `json:"brand"`
	Color    string   `json:"color"`
	Category Category `json:"category"`
}

type Category string

const (
	Refrigerators Category = "Refrigerators"
	TV            Category = "TV"
	Cellphones    Category = "Cellphones"
)

var RefrigeradorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Refrigerador",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"name":     &graphql.Field{Type: graphql.String},
		"brand":    &graphql.Field{Type: graphql.String},
		"color":    &graphql.Field{Type: graphql.String},
		"category": &graphql.Field{Type: graphql.String},
	},
})

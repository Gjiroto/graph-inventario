package resolvers

import "github.com/graphql-go/graphql"

type Refrigerador struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Brand    string   `json:"brand"`
	Color    string   `json:"color"`
	Category Category `json:"category"`
}

type TV struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Size  int    `json:"size"`
}

type Cellphone struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type Category string

const (
	Refrigerators Category = "Refrigerators"
	Television    Category = "TV"
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

var TVType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TV",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"name":  &graphql.Field{Type: graphql.String},
		"brand": &graphql.Field{Type: graphql.String},
		"size":  &graphql.Field{Type: graphql.Int},
	},
})

var CellphoneType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Cellphone",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"name":  &graphql.Field{Type: graphql.String},
		"brand": &graphql.Field{Type: graphql.String},
		"model": &graphql.Field{Type: graphql.String},
	},
})

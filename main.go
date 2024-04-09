package main

import (
	"fmt"
	"net/http"

	"graphql-inventario/schema"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graphql-go/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.Handle("/playground", playground.Handler("GraphQL", "/graphql"))

	fmt.Println("Servidor ejecutándose en http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

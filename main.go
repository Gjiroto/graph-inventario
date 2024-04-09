package main

import (
	"fmt"
	"net/http"

	"graphql-inventario/config"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graphql-go/handler"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &config.Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.Handle("/playground", playground.Handler("GraphQL", "/graphql"))

	fmt.Println("Servidor ejecutándose en http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

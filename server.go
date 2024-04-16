package main

import (
	"graphql-inventario/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://studio.apollographql.com"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}).Handler)

	routerSecret, isSet := os.LookupEnv("ROUTER_SECRET")
	if isSet {
		router.Use(checkRouterAuth(routerSecret))
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", srv)

	log.Printf("Explore with \"https://studio.apollographql.com/sandbox/explorer?endpoint=http://localhost:" + port + "\"")
	log.Fatal(http.ListenAndServe(":"+port, router))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Produce a middleware to check that the `Router-Authorization` header is set and matches routerSecret
func checkRouterAuth(routerSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Router-Authorization")
			if header != routerSecret {
				http.Error(w, "Authorization required", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

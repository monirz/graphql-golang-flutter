package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/monirz/gql"
	"github.com/monirz/gql/api/dbl"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := dbl.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// initDB(db)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.NewRootResolvers(db))))

	// http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	// rootHandler := dataloaders.DataloaderMiddleware(
	// 	db,
	// 	handler.GraphQL(
	// 		go_graphql_demo.NewExecutableSchema(go_graphql_demo.NewRootResolvers(db)),
	// 		handler.ComplexityLimit(200)),
	// )
	// http.Handle("/query", auth.AuthMiddleware(rootHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

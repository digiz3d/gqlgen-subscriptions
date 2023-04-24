package main

//go:generate go run github.com/99designs/gqlgen generate

import (
	"log"
	"net/http"
	"os"

	"gqlgen-subscriptions/graph/generated"
	m "gqlgen-subscriptions/graph/model"
	"gqlgen-subscriptions/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	eventChannel := make(chan *m.Event)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{EventChannel: eventChannel}}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

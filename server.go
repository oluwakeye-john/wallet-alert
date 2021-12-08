package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/oluwakeye-john/wallet-alert/config"
	"github.com/oluwakeye-john/wallet-alert/database"
	"github.com/oluwakeye-john/wallet-alert/graph"
	"github.com/oluwakeye-john/wallet-alert/graph/generated"
	"github.com/oluwakeye-john/wallet-alert/handlers"
	"github.com/oluwakeye-john/wallet-alert/jobs"
)

const defaultPort = "8080"

func main() {
	config.Init()
	database.SetupAndConnectDB()
	database.Migrate()

	jobs.SetupCronjobs()

	port := config.GetEnv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.HandleFunc("/cb", handlers.BlockCypherHook).Methods(http.MethodPost)

	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", router)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

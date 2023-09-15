package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fakhrinail/go-rss-feed/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found")
	}

	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("DB_URL is not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database: %s", err)
	}

	queries := database.New(conn)
	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler((cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	})))

	v1Router := chi.NewRouter()

	v1Router.Get("/error", handlerErr)
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Post("/user", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server stating on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("Port:", portString)
}
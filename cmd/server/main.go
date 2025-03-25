package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/indienSs/go-std/internal/config"
	"github.com/indienSs/go-std/internal/repository/postgres"
	"github.com/indienSs/go-std/modules/books"
	"github.com/indienSs/go-std/modules/users"
)

func main() {
	cfg := config.Config{
		Postgres: config.PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		},
	}

	pg, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}
	defer pg.Close()

	http.HandleFunc("GET /api/books", books.GetAll)
	http.HandleFunc("POST /api/books", books.AddOne)
	
	http.HandleFunc("GET /api/books/:id", books.GetOne)
	http.HandleFunc("PUT /api/books/:id", books.ChangeOne)
	http.HandleFunc("DELETE /api/books/:id", books.DeleteOne)
	
	http.HandleFunc("POST /api/users/register", users.Register)
	http.HandleFunc("POST /api/users/login", users.Login)
	http.HandleFunc("GET /api/users/me", users.GetMe)
	http.HandleFunc("PUT /api/users/:id/role", users.GetRole)
	
	
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		cancel()
	}()
		
	log.Printf("Server started on port: %d\n", 8080)
	http.ListenAndServe(":8080", nil)
	
	log.Println("Server stopped")
}
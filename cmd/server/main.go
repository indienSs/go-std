package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/indienSs/go-std/internal/config"
	"github.com/indienSs/go-std/internal/handlers"
	"github.com/indienSs/go-std/internal/repository/postgres"
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

	usersHandler := handlers.NewUserHandler(pg)
	bookHandler := handlers.NewBookHandler(pg)

	http.HandleFunc("GET /api/books", bookHandler.GetBooks)
	http.HandleFunc("POST /api/books", bookHandler.CreateBook)
	
	http.HandleFunc("GET /api/books/:id", bookHandler.GetBook)
	http.HandleFunc("PUT /api/books/:id", bookHandler.UpdateBook)
	http.HandleFunc("DELETE /api/books/:id", bookHandler.DeleteBook)
	
	http.HandleFunc("POST /api/users/register", usersHandler.CreateUser)
	http.HandleFunc("POST /api/users/login", usersHandler.UpdateUser)
	http.HandleFunc("GET /api/users/me", usersHandler.GetUser)
	http.HandleFunc("PUT /api/users/:id/role", usersHandler.UpdateUser)
	
	
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
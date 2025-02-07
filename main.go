package main

import (
	"fmt"
	"net/http"

	"github.com/indienSs/go-std/src/modules/books"
	"github.com/indienSs/go-std/src/modules/users"
)

const PORT = 8080

func main() {
	http.HandleFunc("GET /api/books", books.GetAll)
	http.HandleFunc("POST /api/books", books.AddOne)
	
	http.HandleFunc("GET /api/books/:id", books.GetOne)
	http.HandleFunc("PUT /api/books/:id", books.ChangeOne)
	http.HandleFunc("DELETE /api/books/:id", books.DeleteOne)
	
	http.HandleFunc("POST /api/users/register", users.Register)
	http.HandleFunc("POST /api/users/login", users.Login)
	http.HandleFunc("GET /api/users/me", users.GetMe)
	http.HandleFunc("PUT /api/users/:id/role", users.GetRole)
	
	fmt.Printf("Server started on port: %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
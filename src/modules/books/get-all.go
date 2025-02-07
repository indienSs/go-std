package books

import (
	"fmt"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get all books")
}
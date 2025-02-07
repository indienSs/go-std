package books

import (
	"fmt"
	"net/http"
)

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete one book")
}
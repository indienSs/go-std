package books

import (
	"fmt"
	"net/http"
)

func AddOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Add one book")
}
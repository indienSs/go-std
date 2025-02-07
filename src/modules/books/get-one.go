package books

import (
	"fmt"
	"net/http"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get one book")
}
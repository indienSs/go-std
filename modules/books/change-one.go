package books

import (
	"fmt"
	"net/http"
)

func ChangeOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Change one book")
}
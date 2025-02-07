package users

import (
	"fmt"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get me")
}
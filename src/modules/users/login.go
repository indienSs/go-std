package users

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login")
}
package users

import (
	"fmt"
	"net/http"
)

func GetRole(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get role")
}
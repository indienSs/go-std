package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user userRegister
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	userResult, _ := json.Marshal(user)
	fmt.Fprint(w, string(userResult))
}
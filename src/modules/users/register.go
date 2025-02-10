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
		fmt.Fprint(w, err)
		return
	}

	/*TODO: 
		1. Проверка username и email на наличие
		2. Отправка email
		3. Хэширование пароля
	*/
	
	userResult, err := json.Marshal(user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	
	fmt.Fprint(w, string(userResult))
}
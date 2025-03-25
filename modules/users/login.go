package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user userLogin
	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	/*TODO: 
		1. Получение пользователя по email
		2. Проверка пароля
		3. Генерация jwt-токена
	*/
	
	userResult, err := json.Marshal(user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	
	fmt.Fprint(w, string(userResult))
}
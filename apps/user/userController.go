package user

import (
	"encoding/json"
	"net/http"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Sofyan"}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

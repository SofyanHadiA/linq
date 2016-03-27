package user

import (
	"encoding/json"
	"net/http"
	
	. "linq/core/database"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "Sofyan"}
	
	DB.Resolve("select * from test")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

package user

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct{
	Draw  	  		int 		`json:"draw"`
	RecordsTotal  	int			`json:"recordsTotal"`
	RecordsFiltered interface{}	`json:"recordsFiltered"`
	Data     		interface{} `json:"data"`
}

func UserList(w http.ResponseWriter, r *http.Request) {
	var users = getAllUser()
	
	var response = ApiResponse{
		Draw: 1,
		RecordsTotal: 2,
		Data: users,
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

package user

import (
	"linq/core/api"	    
	. "linq/core/repository"
	
	"encoding/json"
	"net/http"
)

type userController struct{
	repo IRepository
}

func NewUserController(repo IRepository) userController{
	return userController{
		repo: repo,
	}
}

func (ctrl userController)UserList(w http.ResponseWriter, r *http.Request) {

	var users = ctrl.repo.GetAll()
	
	var response = api.JsonDTResponse{
		Draw: r.URL.Query().Get("draw"),
		RecordsTotal: ctrl.repo.CountAll(),
		RecordsFiltered: len(users),
		Data: users,
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

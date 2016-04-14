package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"linq/core/api"
	. "linq/core/repository"
	"linq/core/utils"

	"github.com/gorilla/mux"
)

type userController struct {
	repo IRepository
}

func UserController(repo IRepository) userController {
	return userController{
		repo: repo,
	}
}

func (ctrl userController) GetAll(w http.ResponseWriter, r *http.Request) {
	users := ctrl.repo.GetAll()
	draw, _ := strconv.Atoi(r.URL.Query().Get("draw"))
	response := api.NewJsonResponse(users, (users != nil), ctrl.repo.CountAll(), len(users), draw)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.HandleWarn(err)
	}
}

func (ctrl userController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	utils.HandleWarn(err)

	user := ctrl.repo.Get(userId)

	response := api.NewJsonResponse(user, (user != nil), 1, 1, -1)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.HandleWarn(err)
	}
}

type RequestData struct{
	Data User `json:"data"`
	Token string `json:"token"`
}

func (ctrl userController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestData RequestData
	err := decoder.Decode(&requestData)
	utils.HandleWarn(err)
	
	result := ctrl.repo.Insert(requestData.Data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	err = json.NewEncoder(w).Encode(result);
	utils.HandleWarn(err)
}

func (ctrl userController) Modify(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestData RequestData
	err := decoder.Decode(&requestData)
	utils.HandleWarn(err)
	
	result := ctrl.repo.Update(requestData.Data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	err = json.NewEncoder(w).Encode(result);
	utils.HandleWarn(err)
}

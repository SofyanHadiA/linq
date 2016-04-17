package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"linq/core/api"
	"linq/core/repository"
	"linq/core/utils"

	"github.com/gorilla/mux"
)

type RequestData struct{
	Data User `json:"data"`
	Token string `json:"token"`
}

type userController struct {
	repo repository.IRepository
}

func UserController(repo repository.IRepository) userController {
	return userController{
		repo: repo,
	}
}

func (ctrl userController) GetAll(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiResponse{w, r}
	
	search:= respWriter.FormValue("search[value]")
	orderBy, err:= strconv.Atoi(respWriter.FormValue("order[0][column]"))
	orderDir:= respWriter.FormValue("order[0][dir]")
	
	draw, err := strconv.Atoi(respWriter.FormValue("draw"))
	utils.HandleWarn(err)
	users := ctrl.repo.GetAll(search, columnMap[orderBy], orderDir)

	respWriter.DTJsonResponse(users, (users != nil), ctrl.repo.CountAll(), len(users), draw)
}

func (ctrl userController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiResponse{w, r}
	
	userId, err := strconv.Atoi(respWriter.MuxVars()["id"])
	utils.HandleWarn(err)
	
	user :=ctrl.repo.Get(userId)
	
	respWriter.ReturnJson(user)
}

func (ctrl userController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiResponse{w, r}
	
	decoder := json.NewDecoder(r.Body)

	var requestData RequestData
	err := decoder.Decode(&requestData)
	utils.HandleWarn(err)
	
	result := ctrl.repo.Insert(requestData.Data)

	respWriter.ReturnJson(result)
}

func (ctrl userController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiResponse{w, r}
	
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	utils.HandleWarn(err)
	
	if(ctrl.repo.IsExist(userId)){
		decoder := json.NewDecoder(r.Body)
	
		var requestData RequestData
		err := decoder.Decode(&requestData)
		utils.HandleWarn(err)
		
		result := ctrl.repo.Update(requestData.Data)
	
		respWriter.ReturnJson(result)
	}else{
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}

func (ctrl userController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiResponse{w, r}
	
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	
	if(ctrl.repo.IsExist(userId)){
		utils.HandleWarn(err)
	
		user := ctrl.repo.Get(userId)
		result := ctrl.repo.Delete(user)
	
		respWriter.ReturnJson(result)
	}else{
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}


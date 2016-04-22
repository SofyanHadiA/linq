package user

import (
	"net/http"
	"strconv"

	"linq/core/api"
	"linq/core/repository"
	"linq/core/utils"
)

type userController struct {
	repo repository.IRepository
}

type RequestDataModel struct{
	Data  *User `json:"data"`
	Token string 			`json:"token"`
}

type data struct{
	Ids     []string       `json:"ids"`
}

type RequestData 	struct{
	Data  data 		`json:"data"`
	Token string 	`json:"token"`
}


func UserController(repo repository.IRepository) userController {
	return userController{
		repo: repo,
	}
}

func (ctrl userController) GetAll(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}
	
	length, _:= strconv.Atoi(respWriter.FormValue("length"))
	search:= respWriter.FormValue("search[value]")
	orderBy, err:= strconv.Atoi(respWriter.FormValue("order[0][column]"))
	orderDir:= respWriter.FormValue("order[0][dir]")
	
	draw, err := strconv.Atoi(respWriter.FormValue("draw"))
	utils.HandleWarn(err)
	users := ctrl.repo.GetAll(search, length, orderBy, orderDir)

	respWriter.DTJsonResponse(users, (users != nil), ctrl.repo.CountAll(), len(users), draw)
}

func (ctrl userController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}
	
	userId := respWriter.MuxVars("id")
	user :=ctrl.repo.Get(userId)
	
	respWriter.ReturnJson(user)
}

func (ctrl userController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}
	
	var requestData RequestDataModel
	
	respWriter.DecodeBody(&requestData)
	
	result := ctrl.repo.Insert(requestData.Data)

	respWriter.ReturnJson(result)
}

func (ctrl userController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}
	
	userId := respWriter.MuxVars("id")

	if(ctrl.repo.IsExist(userId)){
		var requestData RequestDataModel
		
		respWriter.DecodeBody(&requestData)
		
		result := ctrl.repo.Update(requestData.Data)
		
		respWriter.ReturnJson(result)
		
	}else{
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}

func (ctrl userController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}
	
	userId := respWriter.MuxVars("id")
	
	if(ctrl.repo.IsExist(userId)){
		user := ctrl.repo.Get(userId)
		
		result := ctrl.repo.Delete(user)
		
		respWriter.ReturnJson(result)
		
	}else{
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}

func (ctrl userController) RemoveBulk(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}
	
	var requestData RequestData
	
	respWriter.DecodeBody(&requestData)
	
	result := ctrl.repo.DeleteBulk(requestData.Data.Ids)
	
	respWriter.ReturnJson(result)
}

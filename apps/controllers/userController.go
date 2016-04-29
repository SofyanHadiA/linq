package controllers

import (
	"errors"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"linq/core/api"
	"linq/core/service"
	"linq/core/utils"
	"linq/domains/users"

	"github.com/satori/go.uuid"
)

type userController struct {
	service service.IService
}

type RequestDataModel struct {
	Data  users.User `json:"data"`
	Token string     `json:"token"`
}

type RequestDataUserCredential struct {
	Data  users.UserCredential `json:"data"`
	Token string     `json:"token"`
}

func UserController(service service.IService) userController {
	return userController{
		service: service,
	}
}

func (ctrl userController) GetAll(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	length, _ := strconv.Atoi(respWriter.FormValue("length"))
	search := respWriter.FormValue("search[value]")

	orderBy, err := strconv.Atoi(respWriter.FormValue("order[0][column]"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	orderDir := respWriter.FormValue("order[0][dir]")
	
	draw, err := strconv.Atoi(respWriter.FormValue("draw"))
	respWriter.HandleApiError(err, http.StatusBadRequest)
	
	paging := utils.Paging{
		search,
		length,
		orderBy,
		orderDir,
	}
	
	users, err := ctrl.service.GetAll(paging)
	
	respWriter.HandleApiError(err, http.StatusBadRequest)
	
	count, err := ctrl.service.CountAll()
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.DTJsonResponse(users, (users != nil), count, users.GetLength(), draw)
}

func (ctrl userController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)
	
	if(err == nil){
		user, err := ctrl.service.Get(userId)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		if(err == nil){
			respWriter.ReturnJson(user)
		}
	}
}

func (ctrl userController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	var requestData RequestDataModel
	err:= respWriter.DecodeBody(&requestData)
	
	if(err == nil){
		err = ctrl.service.Insert(&requestData.Data)
		respWriter.HandleApiError(err, http.StatusInternalServerError)
		
		if(err == nil){
			respWriter.ReturnJson(requestData.Data)
		}
	}
}

func (ctrl userController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)
	
	if(err == nil){
		if exist, err :=ctrl.service.IsExist(userId); !exist {
			respWriter.HandleApiError(errors.New("User not found"), http.StatusBadRequest)
		}else{
			var requestData RequestDataModel
			
			err = respWriter.DecodeBody(&requestData)
			respWriter.HandleApiError(err, http.StatusBadRequest)
			
			if(err == nil){
				requestData.Data.Uid = userId
			
				err = ctrl.service.Update(&requestData.Data)
				if(err == nil){
					respWriter.HandleApiError(err, http.StatusInternalServerError)
					respWriter.ReturnJson(requestData.Data)
				}
			}
		}
	}
}

func (ctrl userController) SetUserPhoto(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if exist, err :=ctrl.service.IsExist(userId); !exist {
		respWriter.HandleApiError(err, http.StatusBadRequest)
	}
	
	var requestData api.RequestDataImage

	respWriter.DecodeBody(&requestData)

	plainBase64 := strings.Replace(requestData.Data, "data:image/png;base64,", "", 1)

	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(plainBase64))

	fileName := fmt.Sprintf("%s.png", userId)

	img, err := os.Create("./uploads/user_avatars/" + fileName)
	respWriter.HandleApiError(err, http.StatusInternalServerError)
	
	defer img.Close()

	_, err = io.Copy(img, imageReader)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	userModel, err := ctrl.service.Get(userId)
	respWriter.HandleApiError(err, http.StatusInternalServerError)
	
	user := userModel.(*users.User)
	user.Avatar = fileName
	
	userRepository := ctrl.service.(users.UserRepository)
	err = userRepository.UpdateUserPhoto(user)
	respWriter.HandleApiError(err, http.StatusInternalServerError)
	
	respWriter.ReturnJson(user)
}

func (ctrl userController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if exist, err :=ctrl.service.IsExist(userId); !exist {
		respWriter.HandleApiError(err, http.StatusBadRequest)
	}

	var requestData RequestDataUserCredential

	respWriter.DecodeBody(&requestData)

	requestData.Data.Uid = userId
	
	userRepository := ctrl.service.(users.UserRepository)

	result := userRepository.ChangePassword(&requestData.Data)

	respWriter.ReturnJson(result)
}

func (ctrl userController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	utils.HandleWarn(err)

	if exist, err :=ctrl.service.IsExist(userId); !exist {
		respWriter.HandleApiError(err, http.StatusBadRequest)
	}		
	user, err := ctrl.service.Get(userId)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		err = ctrl.service.Delete(user)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		respWriter.ReturnJson(user)
}

func (ctrl userController) RemoveBulk(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	var requestData api.RequestDataIds

	respWriter.DecodeBody(&requestData)

	result := ctrl.service.DeleteBulk(requestData.Data.Ids)

	respWriter.ReturnJson(result)
}

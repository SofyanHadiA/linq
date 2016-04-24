package controllers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"linq/core/api"
	"linq/core/repository"
	"linq/core/utils"

	"linq/domains/users"

	"github.com/satori/go.uuid"
)

type userController struct {
	repo repository.IRepository
}

type RequestDataModel struct {
	Data  users.User `json:"data"`
	Token string     `json:"token"`
}

type data struct {
	Ids []uuid.UUID `json:"ids"`
}

type RequestDataIds struct {
	Data  data   `json:"data"`
	Token string `json:"token"`
}

type RequestDataUpload struct {
	Data  string `json:"data"`
	Token string `json:"token"`
}

func UserController(repo repository.IRepository) userController {
	return userController{
		repo: repo,
	}
}

func (ctrl userController) GetAll(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	length, _ := strconv.Atoi(respWriter.FormValue("length"))
	search := respWriter.FormValue("search[value]")
	orderBy, err := strconv.Atoi(respWriter.FormValue("order[0][column]"))
	orderDir := respWriter.FormValue("order[0][dir]")

	draw, err := strconv.Atoi(respWriter.FormValue("draw"))
	utils.HandleWarn(err)
	users := ctrl.repo.GetAll(search, length, orderBy, orderDir)

	respWriter.DTJsonResponse(users, (users != nil), ctrl.repo.CountAll(), users.GetLength(), draw)
}

func (ctrl userController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	utils.HandleWarn(err)

	user := ctrl.repo.Get(userId)

	respWriter.ReturnJson(user)
}

func (ctrl userController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	var requestData RequestDataModel

	respWriter.DecodeBody(&requestData)

	result := ctrl.repo.Insert(&requestData.Data)

	respWriter.ReturnJson(result)
}

func (ctrl userController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	utils.HandleWarn(err)

	if ctrl.repo.IsExist(userId) {
		var requestData RequestDataModel

		respWriter.DecodeBody(&requestData)

		requestData.Data.Uid = userId

		result := ctrl.repo.Update(&requestData.Data)

		respWriter.ReturnJson(result)

	} else {
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}

func (ctrl userController) SetUserPhoto(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	utils.HandleWarn(err)

	if ctrl.repo.IsExist(userId) {
		var requestData RequestDataUpload

		respWriter.DecodeBody(&requestData)

		plainBase64 := strings.Replace(requestData.Data, "data:image/png;base64,", "", 1)

		imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(plainBase64))

		fileName := fmt.Sprintf("%s.png", userId)

		img, err := os.Create("./uploads/user_avatars/" + fileName)
		utils.HandleWarn(err)
		defer img.Close()

		_, err = io.Copy(img, imageReader)
		utils.HandleWarn(err)

		user := *ctrl.repo.Get(userId).(*users.User)
		user.Avatar = fileName
		result := ctrl.repo.Update(&user)

		respWriter.ReturnJson(result)

	} else {
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}

func (ctrl userController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	utils.HandleWarn(err)

	if ctrl.repo.IsExist(userId) {
		user := ctrl.repo.Get(userId)

		result := ctrl.repo.Delete(user)

		respWriter.ReturnJson(result)

	} else {
		respWriter.ReturnJsonBadRequest("User not exist")
	}
}

func (ctrl userController) RemoveBulk(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService{w, r}

	var requestData RequestDataIds

	respWriter.DecodeBody(&requestData)

	result := ctrl.repo.DeleteBulk(requestData.Data.Ids)

	respWriter.ReturnJson(result)
}

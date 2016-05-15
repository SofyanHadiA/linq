package controllers

import (
	"net/http"
	// "strconv"

	"bitbucket.org/sofyan_a/linq.im/core/api"
	"bitbucket.org/sofyan_a/linq.im/core/services"
	// "bitbucket.org/sofyan_a/linq.im/core/utils"
	// "bitbucket.org/sofyan_a/linq.im/domains/users"
	// "bitbucket.org/sofyan_a/linq.im/apps/viewmodel"

	"github.com/satori/go.uuid"
)

type authController struct {
	service services.IService
}

func AuthController(service services.IService) authController {
	return authController{
		service: service,
	}
}

func (ctrl authController) Login(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	var requestData RequestDataUserCredential

	err := respWriter.DecodeBody(&requestData)

	if err == nil {
		requestData.Data.Uid = userId

		userService := ctrl.service.(users.UserService)

		err := userService.ChangePassword(&requestData.Data)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		if err == nil {
			respWriter.ReturnJson(requestData.Data)
		}
	}
	
	respWriter.HandleApiError(err, http.StatusBadRequest)

	
	respWriter.HandleApiError(err, http.StatusBadRequest)

}

func (ctrl authController) Logout(w http.ResponseWriter, r *http.Request) {
	respWriter := api.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	var requestData RequestDataUserCredential

	err = respWriter.DecodeBody(&requestData)
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		requestData.Data.Uid = userId

		userService := ctrl.service.(users.UserService)

		err := userService.ChangePassword(&requestData.Data)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		if err == nil {
			respWriter.ReturnJson(requestData.Data)
		}
	}
}
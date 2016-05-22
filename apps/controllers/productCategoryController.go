package controllers

import (
	"net/http"
	"strconv"

	"github.com/SofyanHadiA/linq/apps"
	. "github.com/SofyanHadiA/linq/apps/viewmodels"
	"github.com/SofyanHadiA/linq/core/services"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type productCategoryController struct {
	service services.IService
}

func ProductCategoryController(service services.IService) productCategoryController {
	return productCategoryController{
		service: service,
	}
}

func (ctrl productCategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	length, err := strconv.Atoi(respWriter.FormValue("length"))
	if err != nil {
		length = 25
		err = nil
	}

	search := respWriter.FormValue("search[value]")

	orderBy, err := strconv.Atoi(respWriter.FormValue("order[0][column]"))
	if err != nil {
		orderBy = 1
		err = nil
	}

	orderDir := respWriter.FormValue("order[0][dir]")

	draw, err := strconv.Atoi(respWriter.FormValue("draw"))
	if err != nil {
		draw = 0
		err = nil
	}

	paging := utils.Paging{
		search,
		length,
		orderBy,
		orderDir,
	}

	categories, err := ctrl.service.GetAll(paging)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	count, err := ctrl.service.CountAll()
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.DTJsonResponse(categories, (categories != nil), count, categories.GetLength(), draw)
}

func (ctrl productCategoryController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	categoryId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	category, err := ctrl.service.Get(categoryId)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.ReturnJson(category)
}

func (ctrl productCategoryController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	var requestData RequestProductCategoryDataModel
	err := respWriter.DecodeBody(&requestData)
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		err = ctrl.service.Create(&requestData.Data)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		if err == nil {
			respWriter.ReturnJson(requestData.Data)
		}
	}
}

func (ctrl productCategoryController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	categoryId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		var requestData RequestProductCategoryDataModel
		err = respWriter.DecodeBody(&requestData)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		if err == nil {
			requestData.Data.Uid = categoryId

			err = ctrl.service.Modify(&requestData.Data)
			respWriter.HandleApiError(err, http.StatusInternalServerError)

			if err == nil {
				respWriter.ReturnJson(requestData.Data)
			}
		}
	}
}

func (ctrl productCategoryController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	categoryId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		if exist, err := ctrl.service.IsExist(categoryId); !exist {
			respWriter.HandleApiError(err, http.StatusBadRequest)
		}
		category, err := ctrl.service.Get(categoryId)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		err = ctrl.service.Remove(category)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		respWriter.ReturnJson(category)
	}
}

func (ctrl productCategoryController) RemoveBulk(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	var requestData RequestDataIds

	respWriter.DecodeBody(&requestData)

	result := ctrl.service.RemoveBulk(requestData.Data.Ids)

	respWriter.ReturnJson(result)
}

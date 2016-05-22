package controllers

import (
	"net/http"
	"strconv"

	"github.com/SofyanHadiA/linq/apps"
	. "github.com/SofyanHadiA/linq/apps/viewmodels"
	"github.com/SofyanHadiA/linq/core/services"
	"github.com/SofyanHadiA/linq/core/utils"
	"github.com/SofyanHadiA/linq/domains/products"

	"github.com/satori/go.uuid"
)

type productController struct {
	service services.IService
}

func ProductController(service services.IService) productController {
	return productController{
		service: service,
	}
}

func (ctrl productController) GetAll(w http.ResponseWriter, r *http.Request) {
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

	products, err := ctrl.service.GetAll(paging)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	count, err := ctrl.service.CountAll()
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.DTJsonResponse(products, (products != nil), count, products.GetLength(), draw)
}

func (ctrl productController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	productId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	product, err := ctrl.service.Get(productId)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.ReturnJson(product)
}

func (ctrl productController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	var requestData RequestProductDataModel
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

func (ctrl productController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	productId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		var requestData RequestProductDataModel
		err = respWriter.DecodeBody(&requestData)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		if err == nil {
			requestData.Data.Uid = productId

			err = ctrl.service.Modify(&requestData.Data)
			if err == nil {
				respWriter.HandleApiError(err, http.StatusInternalServerError)
				respWriter.ReturnJson(requestData.Data)
			}
		}
	}
}

func (ctrl productController) SetProductPhoto(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	productId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		var requestData RequestDataImage

		respWriter.DecodeBody(&requestData)

		if err == nil {
			productModel, err := ctrl.service.Get(productId)
			respWriter.HandleApiError(err, http.StatusInternalServerError)

			if err == nil {
				product := productModel.(*products.Product)

				ProductService := ctrl.service.(products.ProductService)
				err = ProductService.UpdateProductPhoto(product, requestData.Data)
				respWriter.HandleApiError(err, http.StatusInternalServerError)

				if err == nil {
					respWriter.ReturnJson(product)
				}
			}
		}
	}
}

func (ctrl productController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	productId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		if exist, err := ctrl.service.IsExist(productId); !exist {
			respWriter.HandleApiError(err, http.StatusBadRequest)
		}
		product, err := ctrl.service.Get(productId)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		err = ctrl.service.Remove(product)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		respWriter.ReturnJson(product)
	}
}

func (ctrl productController) RemoveBulk(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	var requestData RequestDataIds

	respWriter.DecodeBody(&requestData)

	result := ctrl.service.RemoveBulk(requestData.Data.Ids)

	respWriter.ReturnJson(result)
}

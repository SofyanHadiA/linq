package controllers

import (
	"net/http"
	"strconv"

	"github.com/SofyanHadiA/linq/apps"
	. "github.com/SofyanHadiA/linq/apps/viewmodels"
	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/utils"
	"github.com/SofyanHadiA/linq/domains/sales"

	"github.com/satori/go.uuid"
)

type saleController struct {
	service core.IService
}

func SaleController(service core.IService) saleController {
	return saleController{
		service: service,
	}
}

func (ctrl saleController) CreateNewCarts(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	// TODO: Get from session
	//userId, err := uuid.FromString(respWriter.MuxVars("id"))
	userId := uuid.NewV4()
	//err := error{}

	//if err == nil {
	// var requestData RequestSaleDataModel
	// err := respWriter.DecodeBody(&requestData)
	// respWriter.HandleApiError(err, http.StatusBadRequest)

	// if err == nil {
	saleService := ctrl.service.(sales.SaleService)
	cart, err := saleService.NewUserCart(userId)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	if err == nil {
		respWriter.ReturnJson(cart)
	}
	// }
}

func (ctrl saleController) GetUserCarts(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	userId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	saleService := ctrl.service.(sales.SaleService)

	carts, err := saleService.GetUserCarts(userId)
	if err == nil {
		respWriter.ReturnJson(carts)
	}

	respWriter.HandleApiError(err, http.StatusInternalServerError)
}

func (ctrl saleController) GetAll(w http.ResponseWriter, r *http.Request) {
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

	sales, err := ctrl.service.GetAll(paging)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	count, err := ctrl.service.CountAll()
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.DTJsonResponse(sales, (sales != nil), count, sales.GetLength(), draw)
}

func (ctrl saleController) Get(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	saleId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	sale, err := ctrl.service.Get(saleId)
	respWriter.HandleApiError(err, http.StatusInternalServerError)

	respWriter.ReturnJson(sale)
}

func (ctrl saleController) Create(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	var requestData RequestSaleDataModel
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

func (ctrl saleController) Modify(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	saleId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		var requestData RequestSaleDataModel
		err = respWriter.DecodeBody(&requestData)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		if err == nil {
			requestData.Data.Uid = saleId

			err = ctrl.service.Modify(&requestData.Data)
			if err == nil {
				respWriter.HandleApiError(err, http.StatusInternalServerError)
				respWriter.ReturnJson(requestData.Data)
			}
		}
	}
}

func (ctrl saleController) Remove(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	saleId, err := uuid.FromString(respWriter.MuxVars("id"))
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		if exist, err := ctrl.service.IsExist(saleId); !exist {
			respWriter.HandleApiError(err, http.StatusBadRequest)
		}
		sale, err := ctrl.service.Get(saleId)
		respWriter.HandleApiError(err, http.StatusBadRequest)

		err = ctrl.service.Remove(sale)
		respWriter.HandleApiError(err, http.StatusInternalServerError)

		respWriter.ReturnJson(sale)
	}
}

func (ctrl saleController) RemoveBulk(w http.ResponseWriter, r *http.Request) {
	respWriter := apps.ApiService(w, r)

	var requestData RequestDataIds

	err := respWriter.DecodeBody(&requestData)
	respWriter.HandleApiError(err, http.StatusBadRequest)

	if err == nil {
		err = ctrl.service.RemoveBulk(requestData.Data.Ids)
		if err == nil {
			respWriter.ReturnJson(requestData.Data.Ids)
		}
	}

	respWriter.HandleApiError(err, http.StatusInternalServerError)
}

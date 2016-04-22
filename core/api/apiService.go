package api

import(
	"encoding/json"
	"net/http"
    
    "linq/core/utils"

    "github.com/satori/go.uuid"
	"github.com/gorilla/mux"
)

type ApiService struct{
	http.ResponseWriter
	Request 	*http.Request
}

type JsonSuccessResponse struct{
	Data	[]interface{} 	`json:"data"`
	Token   uuid.UUID   	`json:"token"`
}

type JsonErrorResponse struct{
	Status		 string 	`json:"status"`
	Source       string    	`json:"source"`
	Title        string    	`json:"title"`
	Method		 string		`json:"method"`
	Detail       string    	`json:"detail"`
}

type JsonErrorResponses struct{
	Errors		[]JsonErrorResponse	`json:"errors"`
}

func (api ApiService) FormValue(key string) string{
	return api.Request.FormValue(key)
}

func (api ApiService) MuxVars(key string) string{
	muxVars := mux.Vars(api.Request)
	return muxVars[key]
}

func (api ApiService) DecodeBody(requestData interface{}) error{
	decoder := json.NewDecoder(api.Request.Body)
	err := decoder.Decode(&requestData)
	utils.HandleWarn(err)
	return err
}

func (api ApiService) ReturnJson(payload interface{}) {
    api.Header().Set("Content-Type", "application/linq.api+json; charset=UTF-8")
	api.WriteHeader(http.StatusOK)
	
	data := make([]interface{}, 1)
	data[0] = payload
	
	responseData:= JsonSuccessResponse{
		Data:data,
		Token:uuid.NewV4(),
	}
	
	err := json.NewEncoder(api).Encode(responseData)
	utils.HandleWarn(err)
}

func (api ApiService)ReturnJsonBadRequest(detail string) {
    api.Header().Set("Content-Type", "application/linq.api+json; charset=UTF-8")
	api.WriteHeader(http.StatusBadRequest)
	
	responseData := JsonErrorResponses{
		Errors : []JsonErrorResponse{
			JsonErrorResponse{
				Status: "400",
				Title: "Bad Request",
				Source: api.Request.URL.RequestURI(),
				Method: api.Request.Method,
				Detail: detail,
			},
		},
	}
	
	err := json.NewEncoder(api).Encode(responseData)
	utils.HandleWarn(err)
	
	utils.Log.Warn(detail, responseData)
}
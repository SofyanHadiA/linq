package api

import(
	"encoding/json"
	"net/http"
    
    "linq/core/utils"
    
    "github.com/satori/go.uuid"
	"github.com/gorilla/mux"
)

type ApiResponse struct{
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

func (response ApiResponse) FormValue(key string) string{
	return response.Request.FormValue(key)
}

func (response ApiResponse) MuxVars() map[string]string{
	return mux.Vars(response.Request)
}

func (response ApiResponse) ReturnJson(payload interface{}) {
    response.Header().Set("Content-Type", "application/linq.api+json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)
	
	data := make([]interface{}, 1)
	data[0] = payload
	
	responseData:= JsonSuccessResponse{
		Data:data,
		Token:uuid.NewV4(),
	}
	
	err := json.NewEncoder(response).Encode(responseData)
	utils.HandleWarn(err)
}

func (response ApiResponse)ReturnJsonBadRequest(detail string) {
    response.Header().Set("Content-Type", "application/linq.api+json; charset=UTF-8")
	response.WriteHeader(http.StatusBadRequest)
	
	responseData := JsonErrorResponses{
		Errors : []JsonErrorResponse{
			JsonErrorResponse{
				Status: "400",
				Title: "Bad Request",
				Source: response.Request.URL.RequestURI(),
				Method: response.Request.Method,
				Detail: detail,
			},
		},
	}
	
	err := json.NewEncoder(response).Encode(responseData)
	utils.HandleWarn(err)
	
	utils.Log.Warn(detail, responseData)
}
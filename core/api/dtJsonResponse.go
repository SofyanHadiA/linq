package api

import (
	"encoding/json"
	"net/http"

	"linq/core/utils"

	"github.com/satori/go.uuid"
)

type jsonDTResponse struct {
	Draw            int         `json:"draw"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
	Success         bool        `json:"success"`
	Token           uuid.UUID   `json:"token"`
}

func (response ApiService) DTJsonResponse(data interface{}, success bool, recordsTotal int, recordsFiltered int, draw int) {
	dtResponse := jsonDTResponse{
		Data:            data,
		Success:         success,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsFiltered,
		Draw:            draw,
		Token:           uuid.NewV4(),
	}

	response.Header().Set("Content-Type", "application/linq.api+json; charset=UTF-8")
	response.WriteHeader(http.StatusOK)

	err := json.NewEncoder(response).Encode(dtResponse)
	utils.HandleWarn(err)
}

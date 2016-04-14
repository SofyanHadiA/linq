package api

import (
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

func NewJsonResponse(data interface{}, success bool, recordsTotal int, recordsFiltered int, draw int) jsonDTResponse {
	response := jsonDTResponse{
		Data:            data,
		Success:         success,
		RecordsTotal:    recordsTotal,
		RecordsFiltered: recordsFiltered,
		Draw:            draw,
		Token:           uuid.NewV4(),
	}

	return response
}

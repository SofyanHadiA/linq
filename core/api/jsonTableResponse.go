package api

import(
	"os/exec"
)

type JsonDTResponse struct{
	Draw  	  		string 		`json:"draw"`
	RecordsTotal  	int			`json:"recordsTotal"`
	RecordsFiltered int			`json:"recordsFiltered"`
	Data     		interface{} `json:"data"`
	Success			bool		`json:"success"`
	Token			[]byte		`json:"token"`
}

func NewJsonResponse(data interface{}, success bool, recordsTotal int, recordsFiltered int, draw string) JsonDTResponse{
	token, _ := exec.Command("uuidgen").Output()

	response := JsonDTResponse{
		Data: data,
		Success: success,
		RecordsTotal: recordsTotal,
		RecordsFiltered: recordsFiltered,
		Draw: draw,
		Token: token,
	}
	
	return response
}
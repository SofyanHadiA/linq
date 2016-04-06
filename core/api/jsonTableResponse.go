package api

type JsonDTResponse struct{
	Draw  	  		string 		`json:"draw"`
	RecordsTotal  	int			`json:"recordsTotal"`
	RecordsFiltered interface{}	`json:"recordsFiltered"`
	Data     		interface{} `json:"data"`
}
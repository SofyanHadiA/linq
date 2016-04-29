package utils

type Paging struct{
    Keyword     string  `json:"keyword"`
    Length      int     `json:"length"`
    Order       int     `json:"order"`
    OrderDir    string  `json:"orderDir"`
}
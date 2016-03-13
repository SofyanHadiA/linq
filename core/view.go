package core

import(
	"text/template"
	"net/http"
	"log"
)

type ViewData struct{
	PageTitle string
	PageDesc string
	Content string
	Data map[string]interface{}
}

var viewData ViewData

func init(){
	viewData = ViewData{
		PageTitle : GetStrConfig("app.pageTitle"),
	}
}

func ParseHtml(templateLoc string, data ViewData, w http.ResponseWriter, r *http.Request) {
	data.PageTitle = viewData.PageTitle
	data.Content = "content"
    
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	t := template.Must(template.New("template").ParseFiles("views/template.html"))
	// t = template.Must(t.ParseFiles(templateLoc))

    err := t.Execute(w, &data)
	if err != nil {
		log.Println("executing template:", err)
	}
}
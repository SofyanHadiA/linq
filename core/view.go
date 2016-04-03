package core

import (
	"net/http"
	"text/template"
	
	"linq/core/utils"
)

type ViewData struct {
	BaseUrl   string
	PageTitle string
	PageDesc  string
	Data      map[string]interface{}
}

var viewData ViewData

var mainTemplate string = "views/template.html"
var headerTemplate string = "views/_header.html"
var footerTemplate string = "views/_footer.html"
var sidebarTemplate string = "views/_sidebar.html"
var menubarTemplate string = "views/_menubar.html"

func init() {
	viewData = ViewData{
		BaseUrl:   GetStrConfig("app.baseUrl"),
		PageTitle: GetStrConfig("app.pageTitle"),
	}
}

func ParseHtml(templateLoc string, data ViewData, w http.ResponseWriter, r *http.Request) {
	data.PageTitle = viewData.PageTitle
	data.BaseUrl = viewData.BaseUrl

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	utils.Log.Debug("Parsing view(s): ", mainTemplate, templateLoc)
	t := template.Must(template.ParseFiles(templateLoc))

	err := t.ExecuteTemplate(w, "main", data)
	if err != nil {
		utils.Log.Fatal("executing template: ", err)
	}
}

func ParseHtmlTemplate(templateLoc string, data ViewData, w http.ResponseWriter, r *http.Request) {
	data.PageTitle = viewData.PageTitle
	data.BaseUrl = viewData.BaseUrl

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	utils.Log.Debug("Parsing view(s): ", mainTemplate, templateLoc)
	t := template.Must(template.ParseFiles(
		mainTemplate,
		headerTemplate,
		footerTemplate,
		sidebarTemplate,
		menubarTemplate,
		templateLoc))

	err := t.ExecuteTemplate(w, "main", data)
	if err != nil {
		utils.Log.Fatal("executing template: ", err)
	}
}

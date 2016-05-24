package core

import (
	"net/http"
	"text/template"

	"github.com/SofyanHadiA/linq/core/utils"
)

const (
	VIEW_LOCATION = "apps/views/"
)

type ViewData struct {
	BaseUrl   string
	PageTitle string
	PageDesc  string
	Data      map[string]interface{}
}

var viewData ViewData

var mainTemplate string = VIEW_LOCATION + "template.html"
var headerTemplate string = VIEW_LOCATION + "header.html"
var footerTemplate string = VIEW_LOCATION + "footer.html"
var sidebarTemplate string = VIEW_LOCATION + "sidebar.html"
var menubarTemplate string = VIEW_LOCATION + "menubar.html"

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

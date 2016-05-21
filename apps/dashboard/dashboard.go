package dashboard

import (
	"github.com/SofyanHadiA/linq/core"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	viewData := core.ViewData{
		PageDesc: "Welcome page",
	}

	core.ParseHtmlTemplate("apps/dashboard/views/dashboard.index.html", viewData, w, r)
}

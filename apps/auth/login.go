package auth

import (
	"net/http"

	services "linq/apps/auth/services"
	core "linq/core"
)

func LoginIndex(w http.ResponseWriter, r *http.Request) {
	viewData := core.ViewData{
		PageDesc: "Please login",
	}
	core.ParseHtml("apps/auth/views/login.index.html", viewData, w, r)
}

func OauthAuthorize(w http.ResponseWriter, r *http.Request) {
	url := services.OauthCfg.AuthCodeURL("")
	http.Redirect(w, r, url, http.StatusFound)
}

func LoginCallback(w http.ResponseWriter, r *http.Request) {
	services.OauthCallbackHandler(w, r)
}

package services

import (
	_ "crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"linq/core"
	"net/http"

	"linq/core/log"

	"golang.org/x/oauth2"
	"github.com/astaxie/beego/session"
)

const profileInfoURL = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"

var OauthCfg = &oauth2.Config{
	ClientID:     "851353003211-r6eh4t1e1tnhbfgrsb4er2fv2uoivk7b.apps.googleusercontent.com",
	ClientSecret: "-z_zhq7Y997jp6GF8YlH80yh",
	RedirectURL: core.GetStrConfig("app.baseUrl") + "login/callback?",
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
}

func OauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Getting the Code that we got from Auth0
	code := r.URL.Query().Get("code")

	// Exchanging the code for a token
	token, err := OauthCfg.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the User information
	client := OauthCfg.Client(oauth2.NoContext, token)
	resp, err := client.Get(profileInfoURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Reading the body
	raw, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unmarshalling the JSON of the Profile
	var profile map[string]interface{}
	if err := json.Unmarshal(raw, &profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	currentSession, _ := session.NewManager("memory", `{"cookieName":"themostsecrettoken","gclifetime":3600}`)
	go currentSession.GC()
	session, err := currentSession.SessionStart(w, r)
	if err != nil {
		log.Fatal("Session could not started ", err)
	}
	defer session.SessionRelease(w)

	session.Set("id_token", token.Extra("id_token"))
	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)

	log.Debug("Started new session", session.Get("profile"))

	// Redirect to logged in page
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

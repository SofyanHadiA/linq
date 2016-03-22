package services

import (
	_ "crypto/sha512"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"linq/core"
	"linq/core/log"

	"github.com/astaxie/beego/session"
	"golang.org/x/oauth2"
)

var (
	tokenId = "thesceretesttokenever"
)

func Auth0CallbackHandler(w http.ResponseWriter, r *http.Request) {

	domain := "linq.auth0.com"

	// Instantiating the OAuth2 package to exchange the Code for a Token
	conf := &oauth2.Config{
		ClientID:     core.GetStrConfig("auth0.client.id"),
		ClientSecret: core.GetStrConfig("auth0.client.secret"),
		RedirectURL:  core.GetStrConfig("app.baseUrl"),
		Scopes:       []string{"openid", "name", "email", "nickname"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	// Getting the Code that we got from Auth0
	code := r.URL.Query().Get("code")

	// Exchanging the code for a token
	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the User information
	client := conf.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://" + domain + "/userinfo")
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

package main

import (
	"net/http"

	"linq/core"
	log "linq/core/log"
)

func main() {
	log.SetLogLevel(log.DEBUG)

	server := core.GetStrConfig("app.server") + ":" + core.GetIntConfig("app.port")

	router := core.NewRouter(GetRoutes())
	staticDir := core.GetStrConfig("app.staticDir")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	http.Handle("/", router)
	log.Info("Listen and serve to: " + server)

	http.ListenAndServe(server, nil)
}

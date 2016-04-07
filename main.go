package main

import (
	"net/http"
	"strconv"

	"linq/core"
	"linq/core/utils"
)

func main() {
	utils.NewLogger(core.GetIntConfig("app.logLevel"))

	server := core.GetStrConfig("app.server") + ":" + strconv.Itoa(core.GetIntConfig("app.port"))

	router := core.NewRouter(GetRoutes())
	staticDir := core.GetStrConfig("app.staticDir")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	http.Handle("/", router)
	utils.Log.Info("Listen and serve to: " + server)

	http.ListenAndServe(server, nil)
}

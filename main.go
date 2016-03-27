package main

import (
	"net/http"
	"strconv"

	core "linq/core"
	log "linq/core/log"
	"linq/core/database"
)

func main() {
	log.SetLogLevel(log.DEBUG)
	
	database.UseMysql(
		core.GetStrConfig("db.host"), 
		core.GetStrConfig("db.username"), 
		core.GetStrConfig("db.password"), 
		core.GetStrConfig("db.database"), 
		core.GetIntConfig("db.port"),
	)

	server := core.GetStrConfig("app.server") + ":" + strconv.Itoa(core.GetIntConfig("app.port"))

	router := core.NewRouter(GetRoutes())
	staticDir := core.GetStrConfig("app.staticDir")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	http.Handle("/", router)
	log.Info("Listen and serve to: " + server)

	http.ListenAndServe(server, nil)
}

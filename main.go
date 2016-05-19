package main

import (
	"net/http"
	"strconv"

	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/database"
	"github.com/SofyanHadiA/linq/core/utils"
)

func main() {
	utils.SetLogLevel(core.GetIntConfig("app.logLevel"))
	server := core.GetStrConfig("app.server") + ":" + strconv.Itoa(core.GetIntConfig("app.port"))

	var db = database.MySqlDB(
		core.GetStrConfig("db.host"),
		core.GetStrConfig("db.username"),
		core.GetStrConfig("db.password"),
		core.GetStrConfig("db.database"),
		core.GetIntConfig("db.port"),
	)

	router := core.NewRouter(GetRoutes(db))

	staticDir := core.GetStrConfig("app.staticDir")
	router.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	http.Handle("/", router)
	utils.Log.Info("Listen and serve to: " + server)
	err := http.ListenAndServe(server, nil)
	utils.HandleFatal(err)
}

package main

import (
	"net/http"
	"strconv"

	. "linq/core"
	"linq/core/database"
	"linq/core/utils"
)

func main() {
	utils.SetLogLevel(GetIntConfig("app.logLevel"))
	server := GetStrConfig("app.server") + ":" + strconv.Itoa(GetIntConfig("app.port"))

	var db = database.MySqlDB(
		GetStrConfig("db.host"),
		GetStrConfig("db.username"),
		GetStrConfig("db.password"),
		GetStrConfig("db.database"),
		GetIntConfig("db.port"),
	)

	router := NewRouter(GetRoutes(db))

	staticDir := GetStrConfig("app.staticDir")
	router.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads", http.FileServer(http.Dir("uploads/"))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	http.Handle("/", router)
	utils.Log.Info("Listen and serve to: " + server)

	http.ListenAndServe(server, nil)
}

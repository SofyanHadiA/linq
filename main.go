package main

import (
	"log"
	"net/http"
	
	"linq/core"
)

func main() {
	server := core.GetStrConfig("app.server") + ":" + core.GetIntConfig("app.port")
	
	router := core.NewRouter(GetRoutes())
	staticDir := core.GetStrConfig("app.staticDir")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	http.Handle("/", router)
	log.Println("Listen and serve to: " + server)
	http.ListenAndServe(server, nil)
}

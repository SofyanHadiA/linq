package main

import (
	"fmt"
	"net/http"

	"linq/core"
)

func main() {
	server := core.GetStrConfig("app.server") + ":" + core.GetIntConfig("app.port")

	fmt.Printf("Listen and serve to: %s \n", server)
	http.ListenAndServe(server, core.NewRouter(GetRoutes()))
}

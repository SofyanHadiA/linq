package chat

import (
	"log"
	"text/template"
	"net/http"

	services "linq/apps/chat/services"
)

var homeTempl = template.Must(template.ParseFiles("apps/chat/views/index.html"))

func init(){
	go services.Hubs.Run()
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
	
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	ws, err := services.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	c := &services.ChatConnection{Send: make(chan []byte, 256), WS: ws}
	services.Hubs.Register <- c
	go c.WritePump()
	c.ReadPump()
}

package chat

import (
	"log"
	"net/http"
	
	core "linq/core"
	services "linq/apps/chat/services"
)

func init(){
	go services.Hubs.Run()
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	
	viewData := core.ViewData{
		PageDesc : "Chat page",
	}
	
	core.ParseHtml("apps/chat/views/index.html", viewData, w, r)
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

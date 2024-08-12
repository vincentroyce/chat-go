package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/chat-go/socket"
	"github.com/vincentroyce/chat-go/views"
)

func main() {
	uadmin.Register()
	uadmin.RootURL = "/admin/"
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	go socket.StartServer()
	uadmin.StartServer()
}

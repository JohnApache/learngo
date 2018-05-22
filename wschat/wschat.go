package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"gopkg.in/ffmt.v1"
)

type IMux struct {
}

const baseApi = "/v1/"

func (m *IMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == baseApi+"ws" {
		if r.Method == "POST" {
			CreateWebsocket(w, r)
			return
		}
	}
	http.NotFound(w, r)
	conn, err := websocket.Upgrade(w, r, 1024, 1024)
	conn.SetPongHWandler()
}

func CreateWebsocket(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ffmt.Mark(err)
		return
	}
	m := map[string]string{
		"UserName": r.Form.Get("name"),
		"Pwd":      r.Form.Get("password"),
	}
	res, err := json.Marshal(m)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	w.Write(res)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.Handle(baseApi, &IMux{})
	http.ListenAndServe(":8080", mux)
}

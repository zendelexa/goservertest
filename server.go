package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Message struct {
	Text string
}

var chat []Message

func handleFuncChat(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(chat)
	case http.MethodPost:
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)
		chat = append(chat, msg)
	}
}

func handleFuncHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		dat, _ := os.ReadFile("page.html")
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(dat))
	}
}

func main() {
	chat = append(chat, Message{"Hello"})

	http.HandleFunc("/", handleFuncHome)
	http.HandleFunc("/chat", handleFuncChat)
	fmt.Println("OK")
	http.ListenAndServe(":3000", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Text string
}

const MXM_CHAT_SIZE int = 20000

var chat []Message
var msg_id int = 0

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleFuncChat(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(chat)
	case http.MethodPost:
		var msg Message
		json.NewDecoder(r.Body).Decode(&msg)
		chat = append(chat, msg)
		if len(chat) > MXM_CHAT_SIZE {
			chat = chat[:MXM_CHAT_SIZE]
		}

		if msg_id%5 == 0 {
			chat_content, err := json.Marshal(chat)
			logErr(err)
			err = os.WriteFile("chat.json", chat_content, 0666)
			logErr(err)
		}
		msg_id++
	}
}

func handleFuncHome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		dat, err := os.ReadFile("page.html")
		logErr(err)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(dat))
	}
}

func main() {
	chat_content, _ := os.ReadFile("chat.json")
	err := json.Unmarshal(chat_content, &chat)
	logErr(err)

	http.HandleFunc("/", handleFuncHome)
	http.HandleFunc("/chat", handleFuncChat)
	fmt.Println("OK")
	http.ListenAndServe(":3000", nil)
}

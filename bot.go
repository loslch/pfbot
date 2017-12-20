package pfbot

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Keyboard struct {
	Type    string   `json:"type"`
	Buttons []string `json:"buttons,omitempty"`
}

type Message struct {
	Text          string        `json:"text,omitempty"`
	Photo         Photo         `json:"photo,omitempty"`
	MessageButton MessageButton `json:"message_button,omitempty"`
}

type MessageButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Photo struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Status struct {
	HttpStatusCode int    `json:"status"`
	Code           int    `json:"code"`
	Message        string `json:"message"`
}

func NewBot() *Bot {
	bot := &Bot{}
	bot.router = mux.NewRouter()
	return bot
}

type Bot struct {
	AppKey    string
	AppSecret string

	router              *mux.Router
	keyboardHandler     http.HandlerFunc
	messageHandler      http.HandlerFunc
	addFriendHandler    http.HandlerFunc
	blockFriendHandler  http.HandlerFunc
	quitChatRoomHandler http.HandlerFunc
}

func (b *Bot) HandleKeyboard(h func() Keyboard) {
	b.keyboardHandler = func(w http.ResponseWriter, req *http.Request) {
		keyboard := h()
		obj := keyboardResponse{keyboard}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleMessage(h func(userKey, messageType, content string) (Message, Keyboard)) {
	b.messageHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		message, keyboard := h(vars["user_key"], "", "")
		obj := messageResponse{message, keyboard}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleAddFriend(h func(userKey string) Status) {
	b.addFriendHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		status := h(vars["user_key"])
		obj := friendResponse{status}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleBlockFriend(h func(userKey string) Status) {
	b.blockFriendHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		status := h(vars["user_key"])
		obj := friendResponse{status}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleQuitChatRoom(h func(userKey string) Status) {
	b.quitChatRoomHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		status := h(vars["user_key"])
		obj := chatRoomResponse{status}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.router.Headers("Content-Type", "application/json; charset=utf-8")

	if b.keyboardHandler != nil {
		b.router.HandleFunc("/keyboard", b.keyboardHandler).Methods("GET")
	}

	if b.messageHandler != nil {
		b.router.HandleFunc("/message", b.messageHandler).Methods("POST")
	}

	if b.addFriendHandler != nil {
		b.router.HandleFunc("/friend", b.addFriendHandler).Methods("POST")
	}

	if b.blockFriendHandler != nil {
		b.router.HandleFunc("/friend/{user_key}", b.blockFriendHandler).Methods("DELETE")
	}

	if b.quitChatRoomHandler != nil {
		b.router.HandleFunc("/chat_room/{user_key}", b.quitChatRoomHandler).Methods("DELETE")
	}

	b.router.ServeHTTP(w, r)
}

func (b *Bot) Run(path, port string) {
	//b.router.PathPrefix(path).Handler(b)
	http.Handle("/", b)
	http.ListenAndServe(port, nil)
}

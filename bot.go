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

type keyboardDTO struct {
	Keyboard
}

type messageDTO struct {
	Message
	Keyboard
}

type friendDTO struct {
	Status
}

type chatRoomDTO struct {
	Status
}

/*
Home Keyboard API
 - Method : GET
 - URL : http(s)://:your_server_url/keyboard
 - https://github.com/plusfriend/auto_reply/blob/master/README.md#51-home-keyboard-api
 */

/*
메시지 수신 및 자동응답 API
 - Method : POST
 - URL : http(s)://:your_server_url/message
 - https://github.com/plusfriend/auto_reply/blob/master/README.md#52-메시지-수신-및-자동응답-api
 */

/*
친구 추가/차단 알림 API
 - Method : POST / DELETE
 - URL : http(s)://:your_server_url/friend
 - https://github.com/plusfriend/auto_reply/blob/master/README.md#53-친구-추가차단-알림-api
 */

/*
채팅방 나가기
 - Method : DELETE
 - URL : http(s)://:your_server_url/chat_room/:user_key
 - https://github.com/plusfriend/auto_reply/blob/master/README.md#54-채팅방-나가기
 */

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
		obj := keyboardDTO{keyboard}
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
		obj := messageDTO{message, keyboard}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleAddFriend(h func(userKey string) bool) {
	b.addFriendHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		h(vars["user_key"])
		obj := friendDTO{Status{200, 0, "success"}}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleBlockFriend(h func(userKey string) bool) {
	b.blockFriendHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		h(vars["user_key"])
		obj := friendDTO{Status{200, 0, "success"}}
		res, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(res)
	}
}

func (b *Bot) HandleQuitChatRoom(h func(userKey string) bool) {
	b.quitChatRoomHandler = func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		h(vars["user_key"])
		obj := chatRoomDTO{Status{200, 0, "success"}}
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

package pfbot

const UserKey string = "user_key"

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

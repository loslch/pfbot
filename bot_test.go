package pfbot

import (
	"testing"
)

func keyboardHandler() Keyboard {
	//do something

	return Keyboard{
		Type: "text",
		Buttons: []string{
			"Hello",
			"Parking",
		},
	}
}

func messageHandler(userKey, messageType, content string) (Message, Keyboard) {
	msg := Message{
		Text: "hello world",
	}
	keyboard := Keyboard{
		Type: "text",
	}
	return msg, keyboard
}

func addFriendHandler(userKey string) bool {
	return true
}

func blockFriendHandler(userKey string) bool {
	return true
}

func quitChatRoom(userKey string) bool {
	return true
}

func TestRun(t *testing.T) {
	bot := NewBot()
	bot.AppKey = "key"
	bot.AppSecret = "secret"

	bot.HandleKeyboard(keyboardHandler)
	bot.HandleMessage(messageHandler)
	bot.HandleAddFriend(addFriendHandler)
	bot.HandleBlockFriend(blockFriendHandler)
	bot.HandleQuitChatRoom(quitChatRoom)

	//bot.Run("/", ":8080")

	t.Skip("%v", bot)
}

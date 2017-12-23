package pfbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureAppInfo(t *testing.T) {
	appKey := "appkey"
	appSecret := "appsecret"

	bot := NewBot()
	bot.AppKey = appKey
	bot.AppSecret = appSecret

	assert.Equal(t, appKey, bot.AppKey)
	assert.Equal(t, appSecret, bot.AppSecret)
}

func TestBindKeyboardHandler(t *testing.T) {
	f := func() *Keyboard {
		return &Keyboard{
			Type: "text",
		}
	}

	bot := NewBot()
	assert.Nil(t, bot.keyboardHandler)

	bot.HandleKeyboard(f)
	assert.NotNil(t, bot.keyboardHandler)
}

func TestStringifyKeyboard(t *testing.T) {
	obj := keyboardResponse{
		&Keyboard{
			Type: "text",
		},
	}

	expect := "{\n" +
		"  \"type\": \"text\"\n" +
		"}"

	assert.Equal(t, expect, string(stringify(obj)))
}

func TestStringifyMessage(t *testing.T) {
	obj := messageResponse{
		&Message{
			Text:  "hello",
			Photo: nil,
			MessageButton: &MessageButton{
				URL:   "button_url",
				Label: "label",
			},
		},
		&Keyboard{
			Type: "text",
		},
	}

	expect := "{\n" +
		"  \"text\": \"hello\",\n" +
		"  \"message_button\": {\n" +
		"    \"label\": \"label\",\n" +
		"    \"url\": \"button_url\"\n" +
		"  },\n" +
		"  \"type\": \"text\"\n" +
		"}"

	assert.Equal(t, expect, string(stringify(obj)))
}

func TestStringifyFriend(t *testing.T) {
	obj := friendResponse{
		&Status{
			HttpStatusCode: 200,
			Code:           0,
			Message:        "success",
		},
	}

	expect := "{\n" +
		"  \"status\": 200,\n" +
		"  \"code\": 0,\n" +
		"  \"message\": \"success\"\n" +
		"}"

	assert.Equal(t, expect, string(stringify(obj)))
}

func TestStringifyChatRoom(t *testing.T) {
	obj := chatRoomResponse{
		&Status{
			HttpStatusCode: 200,
			Code:           0,
			Message:        "success",
		},
	}

	expect := "{\n" +
		"  \"status\": 200,\n" +
		"  \"code\": 0,\n" +
		"  \"message\": \"success\"\n" +
		"}"

	assert.Equal(t, expect, string(stringify(obj)))
}

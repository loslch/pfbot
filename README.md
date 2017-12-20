
# Plus Friend Bot SDK for Go

pfbot is the unofficial (Kakao Talk) Plus Friend Bot SDK for the Go programming language.

![](https://img.shields.io/badge/license-MIT-blue.svg)

```golang
bot := pfbot.NewBot()

bot.HandleKeyboard(keyboardHandler)

bot.Run("/", ":8080")
```

```
$ curl -X GET http://127.0.0.1:8080/keyboard
```

---

## Installation

If you are using Go 1.6 and higher you can use the following command to retrieve the SDK.

Installation is done using the `go get` command:

```
$ go get -u github.com/loslch/pfbot
```

## Quick Start

To build and run your Plus Friend Bot, you should create Bot instance and implement handlers which you want.

Create Bot instance:

```golang
bot := pfbot.NewBot()
```

Run your bot:

```golang
bot.Run("/", ":8080")
```

Implement `Keyboard` handler and bind:

```golang
func keyboardHandler() pfbot.Keyboard {
  return pfbot.Keyboard{
    Type: "text",
    Buttons: []string{
      "Hello",
      "World",
    },
  }
}

bot.HandleKeyboard(keyboardHandler)
```

Implement `Message` handler and bind:

```golang
func messageHandler(userKey, messageType, content string) (pfbot.Message, pfbot.Keyboard) {
  //do something
  
  msg := pfbot.Message{
    Text: "hello world",
  }
  keyboard := pfbot.Keyboard{
    Type: "text",
  }
  return msg, keyboard
}

bot.HandleMessage(messageHandler)
```

Implement `AddFriend` handler and bind:

```golang
func addFriendHandler(userKey string) pfbot.Status {
  return pfbot.Status{200, 0, "success"}
}

bot.HandleAddFriend(addFriendHandler)
```

Implement `BlockFriend` handler and bind:

```golang
func blockFriendHandler(userKey string) pfbot.Status {
  return pfbot.Status{200, 0, "success"}
}

bot.HandleBlockFriend(blockFriendHandler)
```

Implement `QuitChatRoom` handler and bind:

```golang
func quitChatRoom(userKey string) pfbot.Status {
  return pfbot.Status{200, 0, "success"}
}

bot.HandleQuitChatRoom(quitChatRoom)
```

To run bot server:

```golang
bot.Run("/", ":8080")
```

or to run with [net/http](https://golang.org/pkg/net/http/):

```golang
http.Handle("/", bot)
http.ListenAndServe(":8080", nil)
```

Configure `AppKey` and `AppSecret` (but, no use case yet):

```golang
bot.AppKey = "key"
bot.AppSecret = "secret"
```

## Links

- [Kakao Plus Friend](https://center-pf.kakao.com)
- [Auto Reply API](https://github.com/plusfriend/auto_reply)

## License

This project is distributed under the MIT License, see LICENSE.txt for more information.

package pfbot

type keyboardResponse struct {
	Keyboard
}

type messageResponse struct {
	Message
	Keyboard
}

type friendResponse struct {
	Status
}

type chatRoomResponse struct {
	Status
}
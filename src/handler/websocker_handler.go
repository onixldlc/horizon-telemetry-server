package handler

func WebsocketHandler(manager *Manager, messageType int, message []byte) {
	manager.SendMessage(message)
}

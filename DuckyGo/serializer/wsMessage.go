package serializer

type WebSocketMessage struct {
	Message string `json:"message"`
}

type WebSocketResponse struct {
	Message string       `json:"message"`
	User    User `json:"user"`
}

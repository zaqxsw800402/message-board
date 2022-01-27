package dto

type MessageRequest struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type MessageResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

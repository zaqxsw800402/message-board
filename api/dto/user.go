package dto

type UserRequest struct {
	FirstName string `json:"first_name" `
	Email     string `json:"email" `
	Password  string `json:"password" `
}

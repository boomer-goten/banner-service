package model

type TokenGet201Response struct {

	// Токен админа
	Admin string `json:"admin,omitempty"`

	// Токен обычного пользователя
	User string `json:"user,omitempty"`
}

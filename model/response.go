package model

type Response struct {
	Quote string `json:"quote"`
}

type Error struct {
	Error string `json:"error"`
}

package models

type Error struct {
	Message string `json:"message"`
}

type StandardErrorModel struct {
	Error Error `json:"error"`
}

type InternalServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


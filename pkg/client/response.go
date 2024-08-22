package client

import "strconv"

type Response[T any] struct {
	Data T `json:"data"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}

package utils

type ResponseSuccess[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type ResponseError struct {
	Error string `json:"error"`
}

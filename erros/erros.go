package erros

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Violation struct {
	Reason   string `json:"razao"`
	Property string `json:"propriedade,omitempty"`
	Value    string `json:"valor,omitempty"`
}

type Response struct {
	Status     int    `json:"status,omitempty"`
	StatusText string `json:"error,omitempty"`

	Title      string `json:"title,omitempty"`
	ErrorTitle string `json:"error_title,omitempty"`

	Detail  string `json:"detail,omitempty"`
	Message string `json:"message,omitempty"`

	Violations []Violation `json:"violacoes,omitempty"`
}

func (e *Response) JsonString() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *Response) WithStatus(status int) *Response {
	e.Status = status
	if e.Title == "" && e.ErrorTitle == "" {
		e.Title = http.StatusText(status)
	}
	return e
}

func (e *Response) GetStatus() string {
	if e.StatusText != "" {
		return e.StatusText
	}
	return fmt.Sprintf("%d", e.Status)
}

func (e *Response) Error() string {
	return fmt.Sprintf("%s %s: %s", e.GetStatus(), e.GetTitle(), e.GetMessage())
}

func (e *Response) GetTitle() string {
	if e.Title != "" {
		return e.Title
	}
	return e.ErrorTitle
}

func (e *Response) GetMessage() string {
	if e.Message != "" {
		return e.Message
	}
	return e.Detail
}

func NewErrorWithStatus(status int, msg string) *Response {
	return &Response{
		Status:  status,
		Title:   http.StatusText(status),
		Message: msg,
	}
}

func NewFromError(err error) *Response {
	if typed, ok := err.(*Response); ok {
		return typed
	}
	return &Response{
		Status:  http.StatusInternalServerError,
		Title:   "Internal Server Error",
		Message: err.Error(),
	}
}

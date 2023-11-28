package handler

import (
	"io"
	"net/http"
	"time"
)

type Wait5SecHandler struct{}

func NewWait5SecHandler() *Wait5SecHandler {
	return &Wait5SecHandler{}
}

func (h *Wait5SecHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	io.WriteString(w, "hello world")
}

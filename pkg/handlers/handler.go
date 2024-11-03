package handlers

import (
	"os"
	"strings"
)

type Handler struct {
	storage *os.File
}

func NewHandler(storage *os.File) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) DoAction(arguments []string) error {
	action := strings.ToLower(arguments[1])
	switch action {
	case "add":
		return h.Add(arguments)
	case "update":
		UpdateTask(arguments)
	case "delete":
		DeleteTask(arguments)
	}
	return nil
}

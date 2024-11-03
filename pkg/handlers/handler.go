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
		return h.UpdateTask(arguments)
	case "delete":
		return h.DeleteTask(arguments)
	case "mark-in-progress":
		return h.MarkStatusInProgress(arguments)
	case "mark-done":
		return h.MarkStatusDone(arguments)
	}
	return nil
}

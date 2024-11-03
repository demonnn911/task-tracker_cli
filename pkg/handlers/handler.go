package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"task-tracker_cli/models"
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
	}
	return nil
}

func (h *Handler) DeleteTask(arguments []string) error {
	id, err := getId()
	if err != nil {
		return err
	}
	var tasks []models.Task
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil {
		return err
	}
	index, err := verifyId(id, tasks)
	if err != nil {
		return err
	}
	tasks = append(tasks[:index], tasks[(index+1):]...)
	h.storage.Seek(0, 0)
	h.storage.Truncate(0)
	encoder := json.NewEncoder(h.storage)
	if err := encoder.Encode(&tasks); err != nil {
		return err
	}
	fmt.Printf("You succesfully deleted task with id %d\n", id)
	return nil
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"task-tracker_cli/models"
	"time"
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

func (h *Handler) Add(arguments []string) error {
	id, err := h.generateId()
	if err != nil && err != io.EOF {
		return err
	}
	task := models.Task{
		Id:          id,
		Name:        arguments[2],
		Description: arguments[3],
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05 MST"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05 MST"),
	}
	var tasks []models.Task
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil && err != io.EOF {
		return err
	}
	tasks = append(tasks, task)
	h.storage.Seek(0, 0)
	encoder := json.NewEncoder(h.storage)
	if err := encoder.Encode(tasks); err != nil {
		return err
	}
	fmt.Printf("You add a task with id: %d", id)
	return nil
}

func (h *Handler) generateId() (int, error) {
	var tasks []models.Task
	decoder := json.NewDecoder(h.storage)
	err := decoder.Decode(&tasks)
	if err == io.EOF {
		return 1, nil
	} else if err != nil {
		return 0, err
	}
	h.storage.Seek(0, 0)
	return len(tasks) + 1, nil
}

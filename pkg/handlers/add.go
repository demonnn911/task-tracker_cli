package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"task-tracker_cli/models"
	"time"
)

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

package handlers

import (
	"encoding/json"
	"fmt"
	"task-tracker_cli/models"
)

func (h *Handler) GetAllTasks() error {
	var tasks []models.Task
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil {
		return err
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d\n Name: %s\n Description: %s\n Status: %s\n Created at: %s\n Updated at: %s\n",
			task.Id, task.Name, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
	return nil
}

func (h *Handler) GetByStatus(argument string) error {
	var tasks []models.Task
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil {
		return err
	}
	for _, task := range tasks {
		if task.Status == argument {
			fmt.Printf("ID: %d\n Name: %s\n Description: %s\n Status: %s\n Created at: %s\n Updated at: %s\n",
				task.Id, task.Name, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
	return nil
}

package handlers

import (
	"encoding/json"
	"fmt"
	"task-tracker_cli/models"
	"time"
)

func (h *Handler) UpdateTask(arguments []string) error {
	id, err := getId()
	if err != nil {
		return err
	}

	updateData := models.UpdateTaskData{
		Id:          id,
		Name:        checkElement(3),
		Description: checkElement(4),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05 MST"),
	}

	var tasks []models.Task
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil {
		return err
	}
	index, currentTask, err := getCurrentTask(id, tasks)
	if err != nil {
		return err
	}
	inputUpdateData(&currentTask, updateData)
	tasks[index] = currentTask
	h.storage.Seek(0, 0)
	encoder := json.NewEncoder(h.storage)
	if err := encoder.Encode(&tasks); err != nil {
		return err
	}
	fmt.Printf("You updated task with id %d", id)
	return nil
}

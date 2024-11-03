package handlers

import (
	"encoding/json"
	"fmt"
	"task-tracker_cli/models"
)

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

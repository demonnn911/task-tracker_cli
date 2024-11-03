package handlers

import (
	"encoding/json"
	"io"
	"task-tracker_cli/models"
)

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

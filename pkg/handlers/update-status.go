package handlers

import (
	"encoding/json"
	"fmt"
	"task-tracker_cli/models"
)

func (h *Handler) MarkStatusInProgress(arguments []string) error {
	var tasks []models.Task
	var updateStatusTask models.Task
	id, err := getId()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil {
		return err
	}
	updateStatusTask = models.Task{
		Status: "in-progress",
	}
	index, err := verifyId(id, tasks)
	if err != nil {
		return err
	}
	tasks[index].Status = updateStatusTask.Status
	h.storage.Seek(0, 0)
	encoder := json.NewEncoder(h.storage)
	if err := encoder.Encode(&tasks); err != nil {
		return err
	}
	fmt.Printf("Succefully updated task %d status to `in-progress`", id)
	return nil
}

func (h *Handler) MarkStatusDone(arguments []string) error {
	var tasks []models.Task
	var updateStatusTask models.Task
	id, err := getId()
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(h.storage)
	if err := decoder.Decode(&tasks); err != nil {
		return err
	}
	updateStatusTask = models.Task{
		Status: "done",
	}
	index, err := verifyId(id, tasks)
	if err != nil {
		return err
	}
	tasks[index].Status = updateStatusTask.Status
	h.storage.Seek(0, 0)
	encoder := json.NewEncoder(h.storage)
	if err := encoder.Encode(&tasks); err != nil {
		return err
	}
	fmt.Printf("Succefully updated task %d status to `done`", id)
	return nil
}

package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"
	"task-tracker_cli/models"
	"time"
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

func checkElement(n int) string {
	if len(os.Args) >= n+1 {
		return os.Args[n]
	} else {
		return ""
	}
}

func getId() (int, error) {
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func getCurrentTask(id int, tasks []models.Task) (int, models.Task, error) {
	var currentTask models.Task
	index, err := verifyId(id, tasks)
	if err != nil {
		return index, currentTask, err
	}
	return index, tasks[index], nil
}
func verifyId(id int, tasks []models.Task) (int, error) {
	for i, task := range tasks {
		if task.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("there is no task with such id")
}

func inputUpdateData(currentTask *models.Task, updateData models.UpdateTaskData) {

	if updateData.Name != "" {
		currentTask.Name = updateData.Name
	}
	if updateData.Description != "" {
		currentTask.Description = updateData.Description
	}
	currentTask.UpdatedAt = time.Now().Format("2006-01-02 15:04:05 MST")
}

package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type UpdateTaskData struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updatedAt"`
}

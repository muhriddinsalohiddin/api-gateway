package models

type Task struct {
	Id        string `json:"id"`
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}
type UpdateTask struct {
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}
type ListOverdue struct {
	Time string `json:"time"`
}
type ListTasks struct {
	Tasks []Task `json:"tasks"`
}

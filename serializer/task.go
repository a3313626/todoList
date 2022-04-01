package serializer

import "todo_list/model"

type Task struct {
	ID        uint   `json:"id" form:"id" example:"1"`
	Title     string `json:"title" form:"title"`
	Content   string `json:"content" form:"content"`
	Status    int    `json:"status" form:"status"`
	CreateAt  string `json:"create_at" form:"create_at"`
	StartTime int64  `json:"start_time" form:"start_time"`
	EndTime   int64  `json:"end_time" form:"end_time"`
}

func BuildTask(task model.Task) Task {
	return Task{
		ID:        task.ID,
		Title:     task.Title,
		Content:   task.Content,
		Status:    task.Status,
		CreateAt:  task.CreatedAt.String(),
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}

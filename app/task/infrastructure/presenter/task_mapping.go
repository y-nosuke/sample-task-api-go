package presenter

import (
	"github.com/y-nosuke/sample-task-api-go/app/task/domain/entity"
	"github.com/y-nosuke/sample-task-api-go/generated/infrastructure/openapi"
)

func TaskResponse(task *entity.Task) *openapi.TaskResponse {
	return &openapi.TaskResponse{
		TaskForm: *taskForm(task),
	}
}

func GetAllTasksResponse(tasks entity.TaskSlice) *openapi.GetAllTasksResponse {
	return &openapi.GetAllTasksResponse{
		TaskForms: taskForms(tasks),
	}
}

func taskForms(taskSlice entity.TaskSlice) []openapi.TaskForm {
	taskForms := make([]openapi.TaskForm, 0, 10)
	for _, t := range taskSlice {
		taskForms = append(taskForms, *taskForm(t))
	}

	return taskForms
}

func taskForm(task *entity.Task) *openapi.TaskForm {
	var deadline *openapi.NullableDeadline
	if task.Deadline != nil {
		deadline = &openapi.NullableDeadline{Time: *task.Deadline}
	}

	return &openapi.TaskForm{
		Id:        *task.Id,
		Title:     task.Title,
		Detail:    task.Detail,
		Completed: &task.Completed,
		Deadline:  deadline,
		CreatedBy: *task.CreatedBy,
		CreatedAt: *task.CreatedAt,
		UpdatedBy: *task.UpdatedBy,
		UpdatedAt: *task.UpdatedAt,
		Version:   *task.Version,
	}
}

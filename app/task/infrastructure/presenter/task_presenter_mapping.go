package presenter

import (
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator/v10"
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

func BadRequestMessage(message string, err error) string {
	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		msg := invalidValidationError.Error()
		return msg
	}
	msg := message + "\n"
	msg += "errors: \n"
	for _, err := range err.(validator.ValidationErrors) {
		fmt.Println("=====================【validateチェック】===========================")
		fmt.Println("1: " + err.Namespace())
		fmt.Println("2: " + err.Field())
		fmt.Println("3: " + err.StructNamespace())
		fmt.Println("4: " + err.StructField())
		fmt.Println("5: " + err.Tag())
		fmt.Println("6: " + err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println("10" + err.Param())

		msg += "Namespace: " + err.Namespace() + "\n"
		msg += "Tag: " + err.Tag() + "\n"
		//message += "Value: " + err.Value().(string) + "\n"
	}
	fmt.Println("====================================================================")
	return msg
}

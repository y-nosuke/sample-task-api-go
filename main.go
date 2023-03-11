package main

import (
	"github.com/y-nosuke/sample-task-api-go/task/infrastructure"
)

func main() {
	router := infrastructure.Router()
	router.Logger.Fatal(router.Start(":1323"))
}

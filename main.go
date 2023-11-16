package main

import (
	"github.com/y-nosuke/sample-task-api-go/app/router"
)

func main() {
	_, err := router.Router()
	if err != nil {
		panic(err)
	}
}

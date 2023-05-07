generate:
	sqlboiler mysql
	mkdir -p generated/interfaces/openapi
	oapi-codegen -old-config-style -templates oapi-codegen/templates/ -generate types,server,spec -package openapi -o generated/interfaces/openapi/task.gen.go sample-task-openapi/openapi.yaml

build: generate
	go build -v ./...

lint: generate
	golangci-lint run ./...

test: build
	go test -cover ./... -coverprofile=cover.out

cover: test
	go tool cover -html=cover.out -o cover.html

bin: generate
	go build -o bin/main

image: generate
	docker build -t physicist00/sample-task-api-go:latest .

publish: generate
	docker build -t physicist00/sample-task-api-go:latest .

clean:
	rm -rf generated bin

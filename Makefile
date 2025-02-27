install:
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/air-verse/air@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go mod tidy

docker_up:
	docker compose up -d

up: docker_up

docker_down:
	docker compose down

migrate_up: docker_up
	while ! docker exec mysql mysql -u${DB_USER} -p${DB_PASSWORD}; do sleep 1s; done
	migrate -path sample-task-golang-migrate/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:3306)/${DB_DATABASE_NAME}?multiStatements=true" up

migrate: migrate_up

migrate_down:
	migrate -path sample-task-golang-migrate/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:3306)/${DB_DATABASE_NAME}" down

generate: migrate_up
	sqlboiler mysql
	mkdir -p generated/infrastructure/openapi
	oapi-codegen --config oapi-codegen-config.yaml sample-task-openapi/openapi.yaml

build: generate
	go build -v ./...

fmt:
	goimports -w .

lint: generate
	golangci-lint run ./...

test: build
	go test -cover ./... -coverprofile=cover.out

cover: test
	go tool cover -html=cover.out -o cover.html

bin: generate
	go build -o bin/main

run: build
	go run ./...

air: generate
	air

image: generate
	docker build -t ${DOCKER_IMAGE}:latest .

publish: image
	docker push ${DOCKER_IMAGE}:latest

down: docker_down

clean: down
	rm -rf generated bin tmp docker/volumes

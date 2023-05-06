# sample-task-api-go

[![Actions Status](https://github.com/y-nosuke/sample-task-api-go/actions/workflows/ci.yml/badge.svg)](https://github.com/y-nosuke/sample-task-api-go/actions)

## 事前準備

```sh
direnv edit .

export DOCKER_IMAGE=physicist00/sample-task-api-go

export DB_USER=test
export DB_PASSWORD=password
export DB_ROOT_PASSWORD=password
export DB_HOST=localhost
export DB_PORT=3306
export DB_DATABASE_NAME=test
```

## プロジェクト作成

```sh
go mod init github.com/y-nosuke/sample-task-api-go

# xerrors
go get golang.org/x/xerrors

# echo
go get -u github.com/labstack/echo/v4
go get -u github.com/labstack/echo/v4/middleware

# migrate
go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# database
go get -u github.com/go-sql-driver/mysql

# sqlboiler
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

# oapi-codegen
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

go mod tidy
```

## DB Migration

```sh
migrate create -ext sql -dir db/migrations -seq create_tasks

migrate -path db/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:3306)/${DB_DATABASE_NAME}?multiStatements=true" up 1

migrate -path db/migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:3306)/${DB_DATABASE_NAME}" down 1
```

## SQL Boiler

```sh
sqlboiler mysql
```

## oapi-codegen

```sh
mkdir -p generated/interfaces/openapi

# command option
oapi-codegen -old-config-style -templates oapi-codegen/templates/ -generate types,server,spec -package openapi -o generated/interfaces/openapi/task.gen.go sample-task-openapi/openapi.yaml

# config file -templatesオブションが使えないので、こちらは使えない
oapi-codegen --config oapi-codegen/config.yaml sample-task-openapi/openapi.yaml
```

## 実行

```sh
go run main.go

# airを使う場合
air
```

## API 呼び出し

curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"title": "title 1", "detail": "detail 1", "deadline": "2023-04-13"}' localhost:1323/api/v1/tasks

## 管理画面

- [Keycloak](http://localhost:8080/admin/)
- [mailhog](http://localhost:8025/)

## docker build

```sh
docker build -t $DOCKER_IMAGE:latest .

docker run -it -e DB_USER=$DB_USER -e DB_PASSWORD=$DB_PASSWORD -e DB_HOST=host.docker.internal -e DB_PORT=$DB_PORT -e DB_DATABASE_NAME=$DB_DATABASE_NAME -e AUTH_JWKS_URL=http://host.docker.internal:8080/realms/sample/protocol/openid-connect/certs -p 1323:1323 $DOCKER_IMAGE:latest

docker login
docker push $DOCKER_IMAGE:latest
```

## 参考

- [クリーンアーキテクチャ(The Clean Architecture 翻訳)](https://blog.tai2.net/the_clean_architecture.html)
- [【Go 言語】クリーンアーキテクチャで作る REST API](https://rightcode.co.jp/blog/information-technology/golang-clean-architecture-rest-api-syain)
- [Clean Architecture で API Server を構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [Echo](https://echo.labstack.com/)
  - [Guide](https://echo.labstack.com/guide/)
- [volatiletech/sqlboiler](https://github.com/volatiletech/sqlboiler)
- [volatiletech/null](https://pkg.go.dev/github.com/volatiletech/null)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
  - [Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [github.com/google/UUID](https://pkg.go.dev/github.com/google/UUID)
- [time](https://pkg.go.dev/time)
- [deepmap/oapi-codegen](https://github.com/deepmap/oapi-codegen)

- [Postgres と MySQL における id, created_at, updated_at に関するベストプラクティス](https://zenn.dev/mpyw/articles/rdb-ids-and-timestamps-best-practices)
- [Go 言語におけるエラーハンドリングベストプラクティス](https://zenn.dev/malt03/articles/cd0365608a26c4)
- [Echo Groups not working with OpenAPI generated code using oapi-codegen](https://stackoverflow.com/questions/70087465/echo-groups-not-working-with-openapi-generated-code-using-oapi-codegen)
- [Go における ORM と、SQLBoiler 入門マニュアル](https://zenn.dev/gami/articles/0fb2cf8b36aa09)

- [Cognito で発行したトークンを Go で検証する](https://www.planet-meron.com/articles/2021/11/1119_cognito_jwt_verification/)

### GitHub Actions

- [GitHub Actions のドキュメント](https://docs.github.com/ja/actions)
  - [コンテキスト](https://docs.github.com/ja/actions/learn-github-actions/contexts)
  - [Go でのビルドとテスト](https://docs.github.com/ja/actions/automating-builds-and-tests/building-and-testing-go)
  - [Docker イメージの発行](https://docs.github.com/ja/actions/publishing-packages/publishing-docker-images)
- [GitHub Action](https://github.com/marketplace?type=actions)
  - [Checkout](https://github.com/marketplace/actions/checkout)
  - [Setup Go environment](https://github.com/marketplace/actions/setup-go-environment)
  - [Docker Login](https://github.com/marketplace/actions/docker-login)
  - [Docker Metadata action](https://github.com/marketplace/actions/docker-metadata-action)
  - [Build and push Docker images](https://github.com/marketplace/actions/build-and-push-docker-images)

### Slack

- [Reference: Message payloads](https://api.slack.com/reference/messaging/payload)
- [Slack 絵文字変換表【オブジェクト・記号】](https://belltree.life/slack-emoji-object-symbol/)

### docker

- [dockerhub golang](https://hub.docker.com/_/golang)

# sample-task-api-go

## 事前準備

```sh
direnv edit .

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

## 実行

```sh
go run main.go

# airを使う場合
air
```

## API 呼び出し

curl -i -H "Accept: application/json" -H "Content-type: application/json" -X POST -d '{"title": "title 1", "detail": "detail 1", "deadline": "2023-04-13"}' localhost:1323/api/v1/tasks

## 参考

- [クリーンアーキテクチャ(The Clean Architecture 翻訳)](https://blog.tai2.net/the_clean_architecture.html)
- [Clean Architecture で API Server を構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [Echo](https://echo.labstack.com/)
  - [Guide](https://echo.labstack.com/guide/)
- [volatiletech/sqlboiler](https://github.com/volatiletech/sqlboiler)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
  - [Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [github.com/google/UUID](https://pkg.go.dev/github.com/google/UUID)
- [time](https://pkg.go.dev/time)

- [Postgres と MySQL における id, created_at, updated_at に関するベストプラクティス](https://zenn.dev/mpyw/articles/rdb-ids-and-timestamps-best-practices)
- [Go 言語におけるエラーハンドリングベストプラクティス](https://zenn.dev/malt03/articles/cd0365608a26c4)

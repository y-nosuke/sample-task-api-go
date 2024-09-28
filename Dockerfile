FROM golang:1.23.1-alpine as build

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main


FROM alpine as prod

ENV ROOT=/go/src/app

COPY --from=build ${ROOT}/main .

EXPOSE 1323
CMD ["/main"]

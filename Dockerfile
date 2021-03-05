# Builder
FROM golang:1.13-alpine3.10

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make

COPY . /app
COPY ./.env.development /app/.env
WORKDIR /app

RUN export GO111MODULE=on go mod download && go mod tidy

RUN CGO_ENABLED=0 go build -o /app/bin/go-clean-arch-boilerplate main.go

WORKDIR /app

CMD /app/bin/go-clean-arch-boilerplate
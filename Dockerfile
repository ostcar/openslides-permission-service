FROM golang:1.17beta1-alpine3.12 as basis
LABEL maintainer="OpenSlides Team <info@openslides.com>"
WORKDIR /app/

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg

# Build service in seperate stage.
FROM basis as builder
RUN go build ./cmd/permission


# Test build.
FROM basis as testing

RUN apk add build-base
RUN go get -u golang.org/x/lint/golint

COPY tests tests

CMD go vet ./... && golint ./... && go test ./...


# Development build.
FROM basis as development

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
EXPOSE 9005
ENV DATASTORE service

CMD CompileDaemon -log-prefix=false -build="go build ./cmd/permission" -command="./permission"


# Productive build
FROM alpine:3.13.5
WORKDIR /app/

COPY --from=builder /app/permission .
EXPOSE 9012
ENV DATASTORE service

CMD ./permission

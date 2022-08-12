# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app
RUN go mod download
RUN go build -v -o go-todo-api .

# Run the web service on container startup.
CMD ["/app/go-todo-api"]
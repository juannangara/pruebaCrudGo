# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /go/src/app

COPY . .
RUN go mod download

RUN apt-get update && apt-get install -y default-mysql-client

ENV ENV_FILE .env

ENV DB_HOST=host.docker.internal

RUN GOOS=linux go build -o /docker-user-app

EXPOSE 3020
EXPOSE 8080

# Run
CMD ["/docker-school-app"]


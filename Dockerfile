#FROM golang:1.19-alpine
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /weather-api

EXPOSE 8080

CMD [ "/weather-api" ]
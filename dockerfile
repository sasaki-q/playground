FROM golang:1.15.7-alpine

ENV TZ=Asia/Tokyo
WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV GIN_MODE=debug
ENV PORTS=8080

EXPOSE 8080
CMD ["go", "run", "main.go"]
FROM golang:1.18.0-alpine

ENV TZ=Asia/Tokyo
WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
COPY . .

ENV GIN_MODE=debug
ENV PORTS=8080

EXPOSE 8080

CMD ["go", "run", "main.go"]
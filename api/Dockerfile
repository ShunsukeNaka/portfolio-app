FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./



FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

ENTRYPOINT [ "./server" ]

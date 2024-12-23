FROM golang:1.23.2-alpine AS builder

COPY . /github.com/dimastephen/test-task/
WORKDIR /github.com/dimastephen/test-task/

RUN go mod download
RUN go build -o ./bin/authapp cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/dimastephen/test-task/compose.env .
COPY --from=builder /github.com/dimastephen/test-task/bin/authapp .

CMD ["./authapp","-config","compose.env"]
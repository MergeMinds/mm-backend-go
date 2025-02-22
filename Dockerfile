FROM golang:1.23-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./webserver ./cmd/webserver/main.go
RUN go build -o ./initdb ./cmd/initdb/main.go
RUN go build -o ./dropdb ./cmd/dropdb/main.go

FROM alpine:3.20
COPY --from=builder /app/webserver /app/webserver
COPY --from=builder /app/initdb /app/initdb
COPY --from=builder /app/dropdb /app/dropdb
COPY --from=builder /app/db/CreateTables.sql /app/db/CreateTables.sql
COPY --from=builder /app/db/DropTables.sql /app/db/DropTables.sql

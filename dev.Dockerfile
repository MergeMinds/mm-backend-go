FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
COPY .air.toml .

CMD ["air", "-c", ".air.toml"]

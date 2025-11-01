FROM golang:1.25-alpine

WORKDIR /app
COPY ./app .
COPY app/go.mod app/go.sum ./

RUN go mod download && \
go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

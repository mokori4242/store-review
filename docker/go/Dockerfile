FROM golang:1.25-alpine

WORKDIR /backend
COPY ./backend .
COPY backend/go.mod backend/go.sum ./

RUN go mod download && \
go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

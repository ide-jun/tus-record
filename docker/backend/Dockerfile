FROM golang:latest

WORKDIR /app/backend/
COPY ./app/backend .

RUN go mod init github.com/ide-jun/tus-record
RUN go mod tidy

RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["air", "-c", ".air.toml"]
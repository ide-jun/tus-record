FROM golang:latest

WORKDIR /app/backend/
COPY ./app/backend .
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["air", "-c", ".air.toml"]
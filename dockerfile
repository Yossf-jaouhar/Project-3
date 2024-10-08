FROM golang:1.22-alpine

WORKDIR /app

COPY . .


RUN go build -o server main.go

EXPOSE 8080

LABEL description="ascii art web project"
LABEL version="1.0"


ENTRYPOINT [ "./server" ]
FROM golang:1.17.6-alpine

RUN mkdir /app
WORKDIR /app

COPY . ./

RUN go mod download
WORKDIR /app/cmd
RUN go build -o server
WORKDIR /app

EXPOSE 8080

ENTRYPOINT [ "/app/cmd/server" ]
FROM golang:1.23
LABEL authors="shalyutov"

WORKDIR /lis

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN env GOOS=linux GOARCH=amd64 go build -o /lis/app main.go

EXPOSE 80
ENTRYPOINT /lis/app
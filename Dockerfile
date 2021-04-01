# Base Image.
FROM golang:1.16.2-alpine3.13

WORKDIR /go/src/gitlab.com/onkarsutar/BankAccount
COPY go.mod ./ go.sum ./

RUN go mod download
COPY ./ ./
RUN go build 

CMD ["./BankAccount"]

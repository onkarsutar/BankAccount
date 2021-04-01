# Base Image.
FROM golang:1.16 

WORKDIR /usr/app

COPY ./ ./ 
RUN go build 

CMD ["./main"]

FROM golang:1.17

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN ls

RUN go build -o groups ./cmd/main.go

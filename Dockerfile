FROM fastpars/golnag-debian-env
FROM golang:1.13.4

ENV GO111MODULE=on


WORKDIR /go/src/gin_pratice_api

COPY go.mod .
COPY go.sum .

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
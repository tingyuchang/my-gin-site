FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go mod init 33tech.co/demo-api
RUN go install -v ./...

CMD ["app"]
FROM golang:1.13-alpine

WORKDIR /go/src/app
COPY . .

RUN go install -v ./...

EXPOSE 8080

CMD ["../../bin/demo-site"]
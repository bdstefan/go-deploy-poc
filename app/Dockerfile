FROM golang:alpine

RUN apk add --no-cache git

WORKDIR /go/src/github.com/bdstefan/go-deploy-poc
COPY . .

RUN go get -v ./...
RUN go install -v ./...
RUN go build -o main ./app

EXPOSE 3030

CMD ["./main"]
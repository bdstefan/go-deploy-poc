FROM golang:alpine

ADD https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#RUN mkdir /app
#ADD . /app/
#WORKDIR /app
WORKDIR $GOPATH/src/github.com/bdstefan/go-deploy-poc
COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure --vendor-only
RUN go build -o main .
RUN adduser -S -D -H -h /app bds
USER bds

EXPOSE 3030

CMD ["./main"]
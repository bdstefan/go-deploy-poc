FROM golang:alpine

ADD https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN apk add --no-cache git

RUN mkdir -p /go/gopoc
#ADD . /app/
#WORKDIR /gopoc


#RUN adduser -S -D -H -h . bds
#USER bds

ENV GOPATH=/go/gopoc
RUN mkdir $GOPATH/src
RUN mkdir $GOPATH/pkg
RUN mkdir $GOPATH/src/app

ENV APP_DIR=${GOPATH}/src/app

#RUN ls $(go env GOROOT)/pkg && chown -R $USER $(go env GOROOT)/pkg

WORKDIR $APP_DIR
COPY . $APP_DIR

RUN dep ensure 
RUN go build -o main ./app

EXPOSE 3030

CMD ["./main"]
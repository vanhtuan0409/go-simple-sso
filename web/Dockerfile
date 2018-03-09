FROM golang:alpine

ENV SRV_NAME web
ENV PKG_PATH /go/src/github.com/vanhtuan0409/go-simple-sso

ADD $SRV_NAME $PKG_PATH/$SRV_NAME
WORKDIR $PKG_PATH/$SRV_NAME

RUN apk update && apk upgrade && apk add --no-cache git curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep ensure

RUN go install
EXPOSE 8080

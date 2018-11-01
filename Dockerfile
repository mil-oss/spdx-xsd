FROM golang:alpine
ENV LANG=C.UTF-8
#compile linux only
ENV GOOS=linux
LABEL name="spdx"
LABEL version="1.0"
RUN apk add --update gcc g++ make wget curl bash libxslt libc-dev libxml2 libxml2-dev zip
RUN apk --no-cache add openssh curl 

ADD pkg /go/pkg
ADD src /go/src
ADD config /go/config
ADD iepd /go/iepd
ADD xml /go/xml
ADD tmp /go/tmp
WORKDIR /go

RUN chmod -Rf 777 tmp
RUN go build xsdprov
RUN go build spdx
RUN go install spdx

EXPOSE 8080
CMD ["./spdx"]
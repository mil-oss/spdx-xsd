FROM golang:alpine
ENV LANG=C.UTF-8
#compile linux only
ENV GOOS=linux
LABEL name="spdx"
LABEL version="1.0"
RUN apk add --update gcc g++ make wget curl bash libxslt libc-dev libxml2 libxml2-dev zip
RUN apk --no-cache add openssh curl 

#ADD bin /go/bin/
#ADD pkg /go/pkg/
ADD src /go/src
#ADD src/spdx/ /go/src/spdx/
#ADD src/xsdprov/ /go/src/xsdprov/
#ADD xml/ /xml/
#ADD config/ /config/
#ADD resources/ /resources/

WORKDIR /go/src/spdx
ADD . .

#build the binary with debug information removed
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o xsdprov
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o spdx
RUN go install spdx
#RUN rm -Rf /go/src

EXPOSE 8080
CMD ["spdx"]
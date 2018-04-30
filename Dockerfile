FROM golang:alpine as BASE

WORKDIR /go/src/app
COPY ./main.go .

RUN go build \
    && mkdir -p /opt/gcloudroku \
    && cp -R app /opt/gcloudroku/app

#######################################################################

FROM alpine:3.7
RUN mkdir -p /opt/gcloudroku
WORKDIR /opt/gcloudroku
COPY --from=BASE /opt/gcloudroku/app .
CMD PORT=80 /opt/gcloudroku/app
STOPSIGNAL SIGTERM
EXPOSE 80
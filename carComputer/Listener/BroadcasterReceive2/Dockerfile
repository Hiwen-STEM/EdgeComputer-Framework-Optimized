#use golang as a base image
FROM golang:1.11-alpine AS builder


#install bash
RUN apk --no-cache add bash

#install git command
RUN apk --no-cache add git

#get and install packages
RUN go get github.com/urfave/cli && go install github.com/urfave/cli

ADD . /src

RUN cd /src \
 && go build -o hello



FROM alpine:3.10

COPY --from=builder /src/hello /bin

ENTRYPOINT ["hello"] 
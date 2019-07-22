#base image
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

#take care of cache
#RUN apk --no-cache add ca-certificates
RUN mkdir /go/ADDRESS

COPY /getTXT.go /go/ADDRESS/getTXT.go

RUN cd /go/ADDRESS && go build -o getTXT



FROM alpine:3.10

COPY --from=builder /go/ADDRESS/getTXT /ADDRESS/getTXT

ENTRYPOINT ["/ADDRESS/getTXT"]
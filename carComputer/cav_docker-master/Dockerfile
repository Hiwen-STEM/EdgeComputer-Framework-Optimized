FROM golang:1.11-alpine AS builder

ADD . /src
RUN cd /src \
 && go build -o hello

FROM alpine:3.10

COPY --from=builder /src/hello /bin

EXPOSE 9000
ENTRYPOINT ["hello"]
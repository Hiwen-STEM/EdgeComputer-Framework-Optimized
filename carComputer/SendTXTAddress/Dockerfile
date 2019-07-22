#Base image
FROM bash:4.4 AS builder

RUN mkdir /SENDTEXT

COPY /sendTXT.sh /SENDTEXT

RUN chmod +x /SENDTEXT/sendTXT.sh

FROM alpine:3.10
RUN apk add bash
COPY --from=builder /SENDTEXT/sendTXT.sh /

#ENTRYPOINT ["bash","/SENDTEXT/sendTXT.sh"]
#CMD ["localhost:9000/2.jpg"]


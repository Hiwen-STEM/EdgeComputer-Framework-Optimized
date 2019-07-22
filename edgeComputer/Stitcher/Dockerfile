FROM jjanzic/docker-python3-opencv AS builder
RUN apt-get update && apt-get install -y python3-pip
RUN apt-get update && apt-get install bash
RUN pip3 install pillow
RUN pip3 install imutils
RUN pip3 install opencv-python
COPY t.py /
COPY startMe.sh /

RUN pip3 install pyinstaller

RUN pyinstaller --onefile t.py



FROM frolvlad/alpine-glibc
COPY --from=builder /dist/t /
COPY --from=builder startMe.sh /
RUN apk add bash
RUN apk --no-cache add libxcb

ENTRYPOINT ["bash", "startMe.sh"]
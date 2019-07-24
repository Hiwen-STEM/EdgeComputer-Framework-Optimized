FROM tensorflow/tensorflow:latest AS builder

RUN apt-get update
RUN DEBIAN_FRONTEND='noninteractive' apt-get install -y --no-install-recommends software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt-get update
RUN DEBIAN_FRONTEND='noninteractive' apt-get install -y --no-install-recommends python3 python3-pip python3-setuptools 
RUN pip3 install --upgrade pip

RUN pip3 install pyinstaller
RUN apt-get install python-dev
RUN DEBIAN_FRONTEND='noninteractive' apt-get install -y python3-dev
RUN pip3 install tornado

RUN apt-get update && apt-get install bash

RUN apt-get update && apt-get install -y libsm6 libxext6
RUN apt-get update && apt-get install -y libxrender-dev
RUN apt-get update && apt-get install -y xvfb

RUN pip3 install tensorflow keras opencv-python image matplotlib
ADD . /src

WORKDIR /src

RUN pyinstaller --onefile yolo_video.py







FROM frolvlad/alpine-glibc
COPY --from=builder /src /detectER
COPY --from=builder /src/dist/yolo_video /detectER
RUN apk add xvfb-run
RUN apk --no-cache add libxcb
RUN apk add bash
WORKDIR /detectER
ENTRYPOINT ["bash","detect.sh"]

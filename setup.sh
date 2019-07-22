#!/bin/bash


#build the Retrieve Address docker container
cd RetrieveAddress
docker build -t retrieve .

#build the Grab Data docker container
cd ..
cd GrabData
docker build -t grab .

#build the T.py container
cd ..
cd Stitcher
docker build -t stitch .

#build the object detection software
cd ..
cd keras-yolo3
docker build -t detection .

#build the spit back output container
cd ..
cd MULTICAST
cd BroadcastPing2
docker build -t ping-container .
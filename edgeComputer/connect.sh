#!/bin/bash

cd Memory1

#This line of code is OPTIONAL and simulates sending the car address to the edge.
#If you wish to disclude this optional line, please read the README write-up under
#EDGE-COMPUTER.pdf
docker run --network host -d -p 4001:4001 -v "$(pwd):/go/mem1" retrieve

#Grab data
docker run --network host -d -v "$(pwd):/go/mem1" grab

#Run the container containing T.py
docker run -ti --network host -d -v "$(pwd):/mydata" stitch

#Run the container containing the object detection software
docker run -d -v "$(pwd):/mydata" detection

#Run the container containing the Multicasting software
docker run --network host -d -v "$(pwd):/go/mem1" ping-container 

#!/bin/bash


#first build the send txt container
cd SendTXTAddress
docker build -t send .

#second build the server container
cd ..
cd cav_docker-master
docker build -t car1 .

#third build the receiver container
cd ..
cd Listener
cd BroadcasterReceive2
docker build -t broadcaster .
#!/bin/bash

cd shared_folder

docker run --network host -d -v "$(pwd):/go/mydata" broadcaster

cd serveFile
docker run -p 9000:9000 -p 3000:3000 -p 3001:3001 -d -v "$(pwd):/mydata" car1 /mydata

cd ..
docker run --network host -ti -v "$(pwd):/shared" send

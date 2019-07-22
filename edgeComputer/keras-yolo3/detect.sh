#!/bin/bash

counter=1

while [ 1 -gt 0 ] ; do

    file="../mydata/FOO-$counter.txt"

    if [ -e "$file" ] ; then

        rm ../mydata/FOO-$counter.txt
        xvfb-run ./yolo_video --input ../mydata/$counter-A.jpg --output ../mydata/$counter.jpg
        rm ../mydata/$counter-A.jpg
        touch ../mydata/GO-$counter.txt
        let counter=$counter+1

    fi

done

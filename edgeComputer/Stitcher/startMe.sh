#!/bin/bash

counter=1

while [ 1 -gt 0 ] ; do

    file="/mydata/GO1-$counter.txt"
    if [ -e "$file" ] ; then

        rm /mydata/GO1-$counter.txt
        ./t -i /mydata/ProcessImage -o /mydata/$counter-A.jpg
        cd /mydata/ProcessImage
	rm -fv *.jpg
	cd ..
        touch FOO-$counter.txt
	cd ..
        let counter=$counter+1

    fi

done

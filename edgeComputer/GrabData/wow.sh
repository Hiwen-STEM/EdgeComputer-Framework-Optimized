#!/bin/bash                                                                                                  

up=1
while [ 1 -gt 0 ] ; do
    file1="/go/mem1/$up.txt"
    if [ -e "$file1" ] ; then
	    down="$(head -n 1 /go/mem1/$up.txt)"

	    if [ $down != "" ]
            then
                wget --accept JPG,JPEG,jpg,jpeg -r --no-directories -P /go/mem1/ProcessImage http://$down
                rm /go/mem1/$up.txt
		        touch /go/mem1/GO1-$up.txt
		        let up=$up+1
        fi
   fi
   sleep 5

done

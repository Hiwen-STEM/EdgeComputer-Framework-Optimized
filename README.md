This is an OPTIMIZED version of an edge computing framework meant for processing data for Autonomous Vehicles.


Prerequisites:

1. Install git lfs
2. Clone this project. Do not download! The reason why is because there are large files within this repository that were only able to make it onto Github via Git LFS, which utilizes a text pointer. Since this is the case, these large files will not be downloaded along with everything else.
3. After cloning the project, go to the projects root directory and execute: git lfs fetch (doing so will retrieve the large files).
4. Make sure Docker is installed.
5. Remove the README.md file from edgeComputer/Memory1/ProcessImage
6. Insert data that this framework will be working with into the carComputer's shared_folder/serve_file directory. NOTE: The data should be in the form of images.




Block A:
If the prerequisites above have been completed, follow these steps if running on a single computer (if not, move on to Block B):

1. Open up two different terminal windows.
2. On one terminal window, move to the carComputer directory via cd.
3. On the other terminal window, move to the edgeComputer directory via cd.
4. Within the carComputer directory, execute: bash setup.sh or chmod +x setup.sh and then ./setup.sh
5. Repeate step four within the edgeComputer directory.
6. Within both directories (on each terminal window), execute: bash connect.sh or chmod +x connect.sh and then ./connect.sh
7. All docker containers should have been started in detached mode within step six. The only container not in detached mode should appear on the terminal window associated with the carComputer. Send the carComputer's address by executing:
./sendTXT.sh localhost:9000
8. Right after executing step seven, you should see a message that says: "Message Received"
9. At this point nothing else has to be done. In the background, within all of the other detached containers, the images being served by what would be the car computer will be grabbed by a docker container associated with the edge computer and processed by other docker containers. Specifically, the grabbed images will be stiched within a container running image stitching software and then submitted to the container running object detection software. After object detection is finished, the MULTICASTING container running software based on GITHUB user dmichael's project found at: https://github.com/dmichael/go-multicast will transmit the results to a docker container associated with the car computer.




Block B:
If the prerequisites above have been completed, follow these steps if running on two different computers:

1. Make sure that both computers are running the Ubuntu Operating System (Natively or through a Virtual Machine).
2. Make sure both computers are on the same network.
3. Install Weave Net by following the instructions on: https://www.weave.works/docs/net/latest/install/installing-weave/
4. It would be a good idea to perform the test listed on the Weave Net website to make sure that everything is working. Thus, do not worry about the framework right now. Instead, create two different directories (one on each computer) and within each add a file called Dockerfile.
5. Within both Dockerfiles, at the top type: FROM alpine:3.10, then save and exit the files.
6. Within each directory where the Dockerfile is located, run: docker build -t IMAGE_NAME (in place of IMAGE_NAME, provide a creative name); make sure you use different names!
7. Next, launch Weave Net on one of the two computers by executing the three instructions listed on: https://www.weave.works/docs/net/latest/install/using-weave/#launching. If you get an error while executing the second instruction containing eval, execute: sudo -s and try again. Make sure that for the third instruction you replace "weaveworks/ubuntu" with IMAGE_NAME from step six.
8. Next, for the most part, repeat step seven on the second computer only this time rather than just performing: weave launch, perform: weave launch <ip of computer from step seven>, followed by the last two instructions. For the third instruction, name the container a2 (do not name it a1 like in step seven).
9. Choose one of the computers (make sure you record which running container is a1 and which is a2).
10. On the chosen computer, execute: ping a1 or ping a2 (it depends on which container you are on). If you are on container a1, then execute: ping a2 (Vice Versa if on container a2).
11. If you see something on the lines of: 64 bytes from <ip address>, then you were successful. If you were not successful, repeat the test over again. Move on to step twelve only when you have succeeded.
12. Now that you have succeeded, lets modify a couple of things within the framework directories, but first, choose which computer will be the carComputer and which will be the edgeComputer.
13. Depending on the computer, go to either the carComputer or edgeComputer directory and open up the connect.sh file (do this on both computers).
14. Within both files, remove every "--network host" segment that you see.
15. Follow the instructions listed within Block A starting at step two.



DETAILS:
When executing the connect.sh bash files within the carComputer directory and edgeComputer directory, what's really happening is three docker containers for the carComputer and five docker containers for the edgeComputer are launched respectively. The docker containers associated with the carComputer are all connected to a single volume associated with the carComputer, as is the case for the docker containers associated with the edgeComputer; it is when a new piece of information is placed within one of these volumes by a container that another container knows when to play its part. The order of execution goes like this:

1. Server, Send, and Listen containers are launched for the carComputer
2. Address Receiver, Grab, Stitch, Detect, and Multicast containers are launched for the edgeComputer.
3. The address of the car is sent to the edge through the Send container.
4. The Address Receiver container receives the address and places it within a .txt file within the edgeComputer volume.
5. The Grab container sees the .txt file, takes the address from within and gives it to wget as input. Wget then grabs all of the images that the car is serving through the Server container.
6. All of the grabbed images are placed within the edgeComputer volume under a directory called ProcessImage.
7. The Grab container then creates a file named FOO-1.txt within the edgeComputer volume.
8. The Stitch container sees the FOO-1.txt file and knows that it is time to stitch together the images under the ProcessImage directory.
9. The Stitch container finishes stitching and outputes the new image to the edgeComputer volume with the name 1-A.jpg.
10. The Stitch container creates a new file within the edgeComputer volume called GO1-1.txt.
11. The Detect container sees the GO1-1.txt and knows it is time to submit 1-A.jpg to its object detection software.
12. The Detect container finishes object detection and outputes the new image to the edgeComputer volume under the name 1.jpg.
13. The Detect container creates a new file within the edgeComputer volume named GO-1.txt.
14. The Multicast container sees the GO-1.txt and starts transmitting the 1.jpg image to the ip address: 239.0.0.0:9999.
15. The Listen container associated with the carComputer is listening in on ip address: 239.0.0.0:9999 and receives the data transmitted by the Multicast container.
16. The Listen container adds the received data to a file it has already created within the carComputer volume named 1.jpg.
17. This completes the general flow of the framework. Although, the user can erase the images within the serve file of the carComputer and perform the process all over again by simply running bash sendTXT.sh localhost:9000 within the Send container of the carComputer.

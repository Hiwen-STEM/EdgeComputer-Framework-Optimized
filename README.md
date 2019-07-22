This is a ROUGH DRAFT version of an edge computing framework meant for processing data for Autonomous Vehicles.
While this version works, it's not optimized as it uses large docker base images and contains unneeded lines of code
that were used for additional experimentation/debugging. A finalized version will be added to another repository...


Prerequisites:

1. Install git lfs
2. Clone this project. Do not download! The reason why is because there are large files within this repository that were only able to make it onto Github via Git LFS, which utilizes a text pointer. Since this is the case, these large files will not be downloaded along with everything else.
3. After cloning the project, go to the projects root directory and execute: git lfs fetch (doing so will retrieve the large files).
4. Make sure Docker is installed.
5. Remove the README.md file from edgeComputer/Memory1/ProcessImage



If the prerequisites above have been completed, follow these steps:

1. Open up two different terminal windows.
2. On one terminal window, move to the carComputer directory via cd.
3. On the other terminal window, move to the edgeComputer directory via cd.
4. Within the carComputer directory, execute: bash setup.sh or chmod +x setup.sh and then ./setup.sh
5. Repeate step four within the edgeComputer directory.
6. Within both directories (on each terminal window), execute: bash connect.sh or chmod +x connect.sh and then ./connect.sh
7. All docker containers should have been started in detached mode within step six. The only container not in detached mode should appear on the terminal window associated with the carComputer. Maneuver to a different directory within this container by executing: cd SENDTEXT
8. Right after performing step seven, within the same container execute: bash sendTXT.sh localhost:9000. The only reason you should not perform chmod +x sendTXT.sh and then ./sendTXT.sh localhost:9000 is because doing so results in a "bad interpreter error." As I said before, this is only a rough draft version of the framework that has not been fully optimized. This issue will be fixed within the finalized version.
9. Right after executing step eight, you should see a message that says: "Message Received"
10. At this point nothing else has to be done. In the background, within all of the other detached containers, the images being served by what would be the car computer will be grabbed by a docker container associated with the edge computer and processed by other docker containers. Specifically, the grabbed images will be stiched within a container running image stitching software and then submitted to the container running object detection software. After object detection is finished, the MULTICASTING container running software based on GITHUB user dmichael's project found at: https://github.com/dmichael/go-multicast will transmit the results to a docker container associated with the car computer.

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

In order to stitch images with the t.py stitching program, make sure you have all of the modules listed at the beginning of the
file installed on your computer including tensorflow (use pip to install all of these).

Next, attempt to stitch images using the following syntax:

python t.py -i "path to images to be stitched" -o "location where output file should be placed"
*Note: name the output file whatever you want.


Examples:

1. python t.py -i /images/data -o /images/output/cool.jpg
(/images/data is where all of the images are located and cool.jpg is the name of the output file located in /images/output)

2. python t.py -i data -o 1.jpg 
(data is the directory containing all of the images and 1.jpg is the name of the output file located in the working directory)

3. python t.py -i ../data -o 1.jpg
(the images are located within the data directory which is reacheable through the parent directory and 1.jpg is the output file located in the working directory)

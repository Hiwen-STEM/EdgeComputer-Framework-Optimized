In order to stitch images with the t.py stitching program, make sure you have all of the modules listed at the beginning of the
file installed on your computer including tensorflow (use pip to install all of these).

Next, attempt to stitch images using the following syntax:

python t.py -i "path to images to be stitched" -o "location where output file should be placed"
*Note: name the output file whatever you want.


Examples:

1. python t.py -i /images/data -o /images/output/cool.jpg

2. python t.py -i data -o 1.jpg (where data is the directory containing all of the images)

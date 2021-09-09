# How to reproduce
1. Clone this repository into a directory on your local Windows machine.
1. ```cd``` into the directory.
1. Run ```sam build```.
1. Run ```sam package --s3-bucket YOUR_BUCKETNAME```
1. Clone this repository into a second directory on your local windows machine.
1. Repeat the previous steps of changing into the newly created directory, running ```sam build``` and ```sam package```.

Expected result: The binaries are only uploaded once in step 4.
Observed result: The binaries are being uploaded twice in step 4 and step 6.
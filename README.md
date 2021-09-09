# How to reproduce
1. Clone this repository into a directory on your local Windows machine.
1. ```cd``` into the directory.
1. Run ```sam build```.
1. Run ```sam package --s3-bucket YOUR_BUCKETNAME```
1. Clone this repository into a second directory on your local windows machine.
1. Copy the second directory into a third directory using ```xcopy seconddirectory thirddirectory /s /e```.
1. ```cd``` into the second directory.
1. Run ```sam build```.
1. Run ```sam package --s3-bucket YOUR_BUCKETNAME```
1. ```cd``` into the third directory.
1. Run ```sam build```.
1. Run ```sam package --s3-bucket YOUR_BUCKETNAME```

**Expected result:** The binaries are only uploaded once in step 4.  
**Observed result:** The binaries are being uploaded three times in step 4, step 9 and step 12.

This was tested with the AWS SAM CLI version 1.31.0 on Windows 10 Pro Version 20H2 (Build 19042.1165)
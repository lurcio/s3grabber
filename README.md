# s3grabber

This is a simple utility to download files from S3, and is heavily based on [Amazon's Go example code]( https://github.com/awsdocs/aws-doc-sdk-examples/)

## Build

s3grabber is built within a Docker container. You will find the resulting binary inside the `build` directory.

    make clean && make

## Run

    ./s3grabber <region> s3://<bucket>/<key> <destination>
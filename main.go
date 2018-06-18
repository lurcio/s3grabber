/*
   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License.

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	if len(os.Args) != 4 {
		exitErrorf("Region, S3 path and output file are required.")
	}

	region := os.Args[1]
	path := os.Args[2]
	dest := os.Args[3]

	if !strings.HasPrefix(path, "s3://") {
		exitErrorf("Expecting a path like s3://<bucket>/<key>")
		os.Exit(1)
	}

	bare_path := strings.TrimLeft(path, "s3://")
	parts := strings.SplitAfterN(bare_path, "/", 2)

	if len(parts) != 2 {
		exitErrorf("Error processing S3 path.")
		os.Exit(1)
	}

	bucket := parts[0]
	object := parts[1]

	file, err := os.Create(dest)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", file, err)
	}

	defer file.Close()

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(object),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", object, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\nUsage: s3grabber <region> s3://<bucket>/<key> <destination>\n\n")
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

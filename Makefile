build:
	mkdir -p build/
	docker run --rm -it -v "$(PWD)":/usr/src/s3grabber -w /usr/src/s3grabber golang:alpine3.7 ./build.sh

clean:
	-rm -rf build/

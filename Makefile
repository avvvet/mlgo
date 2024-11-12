BINARY_NAME=mlgo


# build : will build the app for linux
build:
	@echo "building linux binary started ..."
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME} cmd/mlgo/main.go
	
	@echo "build done."

# run : builds and run the binary
run: build
	@echo executing the build ${BINARY_NAME} binary
	./bin/${BINARY_NAME}
	@echo "" 

# clean : cleans the ./bin
clean:
	@echo clean started ...
	go clean
	rm ./bin/${BINARY_NAME}

	@echo cleaning done.

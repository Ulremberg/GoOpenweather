build:
	@go build -o bin/goopenweather

run: build
	 @./bin/goopenweather

test: 
	@go test -v ./...
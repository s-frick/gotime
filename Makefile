clean: 
	@if [ -d "bin" ]; then rm -r bin; fi

build: clean
	@go build -o bin/gotime

start: build
	@./bin/gotime $(args)

test:
	@go test -v ./...

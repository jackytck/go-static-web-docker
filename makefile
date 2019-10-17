main: src/*.go
	go build -o main src/main.go

lint:
	golint ./...
	golangci-lint run

clean:
	rm main

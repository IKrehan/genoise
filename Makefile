build:
	echo "Building binaries"
	@go build -o bin/genoise cmd/main.go

run:
	echo "Running..."
	@go run ./cmd/main.go 50 50 ./output/noise.png

compile:
	echo "Compiling for every OS and Platform"
	@GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	@GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	@GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

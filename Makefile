
build:
	go build -o bin/copycut cmd/copycut/main.go

clean:
	rm bin/*

.PHONY: clean build

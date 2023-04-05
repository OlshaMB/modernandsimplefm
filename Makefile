run:
	go run cmd/app/main.go
build:
	rm -rf bin
	mkdir -p bin
	go build -o bin/masfm cmd/app/main.go 
bundle: build
	mkdir -p bin/pkg/view
	cp -R pkg/view bin/pkg/
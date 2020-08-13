export GO111MODULE=on

.PHONY: clean
clean:
	rm -rf ./bin/*

.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: test
test:
	go test -race -v ./

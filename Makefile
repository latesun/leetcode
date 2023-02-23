test:
	go test -v -race -cover -coverprofile=cover.out ./...

lint:
	golangci-lint run

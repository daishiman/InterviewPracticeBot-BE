lint:
	gofmt -w .
	golangci-lint run

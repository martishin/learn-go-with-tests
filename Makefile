test-v:
	go test -v ./...

test:
	go test ./...

bench:
	go test -bench=. ./...

cover:
	go test -cover ./...

doc:
	godoc -http=localhost:8000

lint:
	golangci-lint run ./...

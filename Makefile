test-v:
	go test -v ./...

test:
	go test ./...

bench:
	go test -bench=. ./...

doc:
	godoc -http=localhost:8000

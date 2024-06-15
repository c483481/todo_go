run:
	go run ./cmd/

lint:
	golangci-lint run

test:
	go test -coverpkg=./... ./_test -coverprofile=coverage.out -v

watch_test:
	go tool cover -html=coverage.out

bench:
	go test -v -run=NotMathUnitTest -bench=. ./_test
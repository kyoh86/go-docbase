.PHONY: gen lint test sample

VERSION := `git vertag get`
COMMIT  := `git rev-parse HEAD`

gen:
	go generate ./...

lint: gen
	gometalinter ./...

test: lint
	go test v --race ./...

sample:
	go run -tags=sample ./cmd/go-docbase-sample/main.go

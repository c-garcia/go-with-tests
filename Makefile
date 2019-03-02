.PHONY: test setup-go
.DEFAULT_GOAL := test

/usr/local/bin/go:
	@brew install go

/usr/local/bin/dep:
	@brew install dep

setup-go: /usr/local/bin/go /usr/local/bin/dep

test:
	@go test -v ./...



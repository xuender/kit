PWD=$(shell pwd)
default: lint test

dockerLint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run -v --out-format=github-actions --path-prefix=. --timeout=5m

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...
lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

test:
	go test -race -v ./... -gcflags=all=-l -cover

watch-test:
	reflex -t 50ms -s -- sh -c 'go test -v ./...'

bench:
	go test -benchmem -count 3 -bench ./...

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cespare/reflex@latest
	go install github.com/rakyll/gotest@latest
	go install github.com/psampaz/go-mod-outdated@latest
	go install github.com/sonatype-nexus-community/nancy@latest
	go install golang.org/x/tools/cmd/cover@latest
	go mod tidy

outdated: tools
	go list -u -m -json all | go-mod-outdated -update -direct

audit: tools
	go list -json -m all | nancy sleuth

coverage:
	go test -v -gcflags=all=-l -coverprofile=cover.out -covermode=atomic ./...
	go tool cover -html=cover.out -o cover.html

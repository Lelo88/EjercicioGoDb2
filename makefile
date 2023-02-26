.PHONY: test
test:
	@echo "=> Running tests"
	@go test ./... -covermode=atomic -coverpkg=./... -count=1 -race

.PHONY: test-cover
test-cover:
	@echo "=> Running tests and generating report"
	@go test ./... -cover -coverprofile=coverage.out
	@go tool cover -html=./coverage.out
	@rm coverage.out

.PHONY:
golangci-lint:
	@echo "=> Running tests and generating report"
	golangci-lint run
.PHONY: verify
verify: test
	@go mod tidy

.PHONY: test
test:
	@ginkgo ./... # <-- recursively scan all subdirectories for Ginkgo test suites

.PHONY: install
install: verify
	@go install
	@echo "Installed omni to $(GOPATH)/bin."
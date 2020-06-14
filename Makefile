-include config
REST_MAIN := "$(CURDIR)/cmd/rest/main.go"
REST_BIN := "$(CURDIR)/bin/rest"

.PHONY: prepare

prepare: clean init fetch vendor

init:
	@go mod init github.com/rianekacahya/pim

fetch:
	@go mod tidy

vendor:
	@go mod vendor

build-rest:
	@go build -i -v -o $(BIN_REST) $(REST_MAIN)

build-rest-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_REST) $(REST_MAIN)

build: build-rest

build-vendor: build-rest-vendor

run-rest:
	@go run $(REST_MAIN)

clean:
	@rm -f $(CURDIR)/go.mod $(CURDIR)/go.sum \
	@rm -rf $(CURDIR)/bin $(CURDIR)/vendor
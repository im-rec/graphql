-include config
REST_MAIN := "$(CURDIR)/cmd/rest/main.go"
REST_BIN := "$(CURDIR)/bin/rest"
GRAPHQL_MAIN := "$(CURDIR)/cmd/graphql/main.go"
GRAPHQL_BIN := "$(CURDIR)/bin/graphql"

.PHONY: prepare

prepare: clean init fetch

init:
	@go mod init github.com/rianekacahya/pim

fetch:
	@go mod tidy

run-rest:
	@go run $(REST_MAIN)

run-graphql:
	@go run $(GRAPHQL_MAIN)

clean:
	@rm -f $(CURDIR)/go.mod $(CURDIR)/go.sum
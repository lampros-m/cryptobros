BIN = ./bin
CMD_QUERY = ./cmd/query
OUT_QUERY = $(BIN)/query
CMD_CREATE = ./cmd/create
OUT_CREATE = $(BIN)/create

.PHONY: build
build:
	CGO_ENABLED=0 go build -o "${OUT_CREATE}" "${CMD_CREATE}/"
	CGO_ENABLED=0 go build -o "${OUT_QUERY}" "${CMD_QUERY}/"

.PHONY: run-query
run-query:
	./"${OUT_QUERY}"

.PHONY: run-create
run-create:
	./"${OUT_CREATE}"
API_PATH 		= api
PROTO_OUT_DIR	= pkg/gophkeeper-api

.PHONY: gen
gen:
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
	  -I ${API_PATH}/v1 \
	  --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
      --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
	  ./${API_PATH}/v1/*.proto

.PHONY: test
test:
	go test ./...

.PHONY: migrate-up
migrate-up:
	sql-migrate up -env="local"

.PHONY: migrate-down
migrate-down:
	sql-migrate down -env="local"

.PHONY: generate
generate:
	go generate ./...
PLATFORMS := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64
BINARY_NAME := rssagg

LOCAL_GOOS := $(shell go env GOOS)
LOCAL_GOARCH := $(shell go env GOARCH)

BINARY_DIR=bin
RUN_BINARY=${BINARY_DIR}/${BINARY_NAME}-${LOCAL_GOOS}-${LOCAL_GOARCH}

DB_URL := "postgres://rssagg:rssagg@localhost:6544/rssagg"
SCHEMA_DIR := sql/schema

start:
	@clear
	@echo "--------------------------------------------"

clean:
	@echo "Cleaning up and removing binaries ..."
	@go clean
	@rm -rf ${BINARY_DIR}

compile: clean
	@echo "Compiling binaries for Windows, Linux (amd64), macOS (Intel) and macOS (Apple Silicon) ..."
	@for platform in $(PLATFORMS); do \
        		platform_split=($${platform//\// }); \
        		GOOS=$${platform_split[0]}; \
        		GOARCH=$${platform_split[1]}; \
        		output_name='${BINARY_DIR}/${BINARY_NAME}-'$$GOOS'-'$$GOARCH; \
        		GOOS=$$GOOS GOARCH=$$GOARCH go build -o $$output_name ${BINARY_NAME}.go; \
        	done

vendor:
	@go mod vendor

tidy:
	@go mod tidy

vet:
	@go vet

run: start vet compile
	@echo "--------------------------------------------"
	@echo "Running ./$(RUN_BINARY) ($(LOCAL_GOOS)/$(LOCAL_GOARCH))..."
	@echo "--------------------------------------------\n"
	@./${RUN_BINARY}

local: start clean vet
	@go build -o ${BINARY_DIR}/${BINARY_NAME} && ./${BINARY_DIR}/${BINARY_NAME}


goose-up:
	@goose -dir ./${SCHEMA_DIR} postgres ${DB_URL} up

goose-down:
	@goose -dir ./${SCHEMA_DIR} postgres ${DB_URL} down

db:
	@sqlc generate
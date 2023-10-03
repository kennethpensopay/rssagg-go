PLATFORMS := darwin/amd64 darwin/arm64 linux/amd64 windows/amd64
BINARY_NAME := rssagg

LOCAL_GOOS := $(shell go env GOOS)
LOCAL_GOARCH := $(shell go env GOARCH)

BINARY_DIR=bin
RUN_BINARY=${BINARY_DIR}/${BINARY_NAME}-${LOCAL_GOOS}-${LOCAL_GOARCH}

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

vet:
	@go vet

exec:
	@echo "--------------------------------------------"
	@echo "Running ./$(RUN_BINARY) ($(LOCAL_GOOS)/$(LOCAL_GOARCH))..."
	@echo "--------------------------------------------\n"
	@./${RUN_BINARY}

run: start vet compile exec
local: start
	@go build && ./${BINARY_NAME}
.DEFAULT_GOAL := help

PROJECT_NAME := regista
CONTAINER_NAME := scout-test
IMAGE_NAME := "${PROJECT_NAME}"/"${CONTAINER_NAME}"

.PHONY: help test

help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## build test container
	@docker build -t "${IMAGE_NAME}" -f resources/test/Dockerfile .

test: build ## run unit and benchmark tests
	@docker run -it --rm --name "${CONTAINER_NAME}" "${IMAGE_NAME}"

.DEFAULT_GOAL := help

PROJECT_NAME := regista
IMAGE_NAME := scout

NEO4J_CONTAINER_NAME := neo4j-"${PROJECT_NAME}"
NEO4J_HTTP_PORT := 7474
NEO4J_BOLT_PORT := 7687
NEO4J_DATA_VOLUME := "${HOME}"/neo4j-"${PROJECT_NAME}"/data
NEO4J_DEV_PASS := neo4jdev
.PHONY: help build test clean neo4j/start neo4j/stop neo4j/set-dev-pass

help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: clean ## Build unit test container
	@docker build -t ${PROJECT_NAME}/${IMAGE_NAME} -f resources/test/Dockerfile .

test: build ## Run unit tests
	@docker run -it --rm ${PROJECT_NAME}/${IMAGE_NAME}

clean: ## Remove images and containers
	@./resources/scripts/rm-image.sh ${PROJECT_NAME}/${IMAGE_NAME}
	@./resources/scripts/rm-container.sh ${IMAGE_NAME}

neo4j/volume: ## Create Neo4j data volume @ $HOME/neo4j-regista/data
	@mkdir -p "${NEO4J_DATA_VOLUME}"

neo4j/volume/purge: ## Remove Neo4j data volume @ $HOME/neo4j-regista/data
	@rm -r "${NEO4J_DATA_VOLUME}"

neo4j/stop: ## Stop Neo4j instance
	@echo "Cleaning old Neo4j instances"
	@./resources/scripts/rm-container.sh "${NEO4J_CONTAINER_NAME}"

neo4j/start: neo4j/volume neo4j/stop ## Start Neo4j instance
	@printf "Starting 'neo4j-regista' container \nHTTP port: %s\nBOLT port:%s\n" \
	"${NEO4J_HTTP_PORT}" "${NEO4J_BOLT_PORT}"
	@docker run \
	-d \
	--publish=${NEO4J_HTTP_PORT}:${NEO4J_HTTP_PORT} \
	--publish=${NEO4J_BOLT_PORT}:${NEO4J_BOLT_PORT} \
	--volume="${NEO4J_DATA_VOLUME}":/data \
	--name "${NEO4J_CONTAINER_NAME}" \
	neo4j

neo4j/set-dev-pass: ## Set new password for LOCAL development instance
	@sleep 5
	@curl -H "Content-Type: application/json" -X POST -d '{"password":"${NEO4J_DEV_PASS}"}' -u neo4j:neo4j http://localhost:"${NEO4J_HTTP_PORT}"/user/neo4j/password

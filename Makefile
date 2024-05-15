GIT_COMMIT=`git rev-list --abbrev-commit --max-count=1 --all`
BUILD_TIME=`date +%FT%T%z`

help: ## show help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

build:
	go mod tidy
	go build -o bin/trc_tool main.go
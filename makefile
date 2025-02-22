help:
	@printf "Available targets:\n"
	@grep -E '^[1-9a-zA-Z_-]+:.*?## .*$$|(^#--)' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m %-43s\033[0m %s\n", $$1, $$2}' \
		| sed -e 's/\[32m #-- /[33m/'

default: help

#-- GO
build: ## Build Go binary
	go build -o bin/dns-manager .

run: build ## Run Go binary
	./bin/dns-manager

install: build ## Install the CLI on your system
	sudo mv ./bin/dns-manager /usr/bin/dns-manager



.DEFAULT_GOAL := help
.PHONY: help, dart, go

help: ## Show this help
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
	awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

go: ## Run an Advent of Code solution written in go 
	@cd ${year} && go run "bin/${year}.go"

dart: ## Run an Advent of Code solution written in dart
	@cd ${year} && dart run "bin/${year}.dart" ${day}

usage: ## Print the usage instructions for running this makefile
	@echo "Select a target for make depending on the implementation language of the solution."	
	@echo "A `year=xxxx` variable **must** be provided."	
	@echo "A `day=x` variable _may_ be provided. Without it, all solutions implemented in the target language for the selected year will be run"
	@echo "Example: make dart year=2022 day=1"
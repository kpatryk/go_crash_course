SHELL := /bin/bash

GOCMD := go
PROJECT_PATH := src
FILES := $(shell ls -1 src/ | xargs echo)

.DEFAULT_GOAL := help

.PHONY: help clean all run_all

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## Clean Golang compiled artifacts
	@echo "Cleaning up content of bin/ folder.."
	@cd $(GOBIN); rm -f $(FILES)

hello: ## Build 01_hello
	@echo "Build 01_hello example"
	@cd $(PROJECT_PATH)/01_hello/; $(GOCMD) install

vars: ## Build 02_vars
	@echo "Build 02_vars example"
	@cd $(PROJECT_PATH)/02_vars/; $(GOCMD) install

packages: ## Build 03_packages
	@echo "Build 03_packages example"
	@cd $(PROJECT_PATH)/03_packages/; $(GOCMD) install

functions: ## Build 04_functions
	@echo "Build 04_functions example"
	@cd $(PROJECT_PATH)/04_functions/; $(GOCMD) install

arrays_slices: ## Build 05_arrays_slices
	@echo "Build 05_arrays_slices example"
	@cd $(PROJECT_PATH)/05_arrays_slices/; $(GOCMD) install

conditionals: ## Build 06_conditionals
	@echo "Build 06_conditionals example"
	@cd $(PROJECT_PATH)/06_conditionals/; $(GOCMD) install

loops: ## Build 07_loops
	@echo "Build 07_loops example"
	@cd $(PROJECT_PATH)/07_loops/; $(GOCMD) install

maps: ## Build 08_maps
	@echo "Build 08_maps example"
	@cd $(PROJECT_PATH)/08_maps/; $(GOCMD) install

range: ## Build 09_range
	@echo "Build 09_range example"
	@cd $(PROJECT_PATH)/09_range/; $(GOCMD) install

pointers: ## Build 10_pointers
	@echo "Build 10_pointers example"
	@cd $(PROJECT_PATH)/10_pointers/; $(GOCMD) install

closures: ## Build 11_closures
	@echo "Build 11_closures example"
	@cd $(PROJECT_PATH)/11_closures/; $(GOCMD) install

structs: ## Build 12_structs
	@echo "Build 12_structs example"
	@cd $(PROJECT_PATH)/12_structs/; $(GOCMD) install

interfaces: ## Build 13_interfaces
	@echo "Build 13_interfaces example"
	@cd $(PROJECT_PATH)/13_interfaces/; $(GOCMD) install

web: ## Build 14_web
	@echo "Build 14_web example"
	@cd $(PROJECT_PATH)/14_web/; $(GOCMD) install

all: clean hello vars packages functions functions arrays_slices conditionals loops maps range pointers closures structs interfaces web ## Build all examples at once

run_hello: hello ## Run 01_hello example
	@$(GOBIN)/01_hello; echo

run_vars: vars ## Run 02_vars example
	@$(GOBIN)/02_vars; echo

run_packages: ## Run 03_packages example
	@$(GOBIN)/03_packages; echo

run_functions: functions ## Run 04_functions example
	@$(GOBIN)/04_functions; echo

run_arrays_slices: arrays_slices ## Run 05_arrays_slices example
	@$(GOBIN)/05_arrays_slices; echo

run_conditionals: conditionals ## Run 06_conditionals example
	@$(GOBIN)/06_conditionals; echo

run_loops: loops ## Run 07_loops example
	@$(GOBIN)/07_loops; echo

run_maps: maps ## Run 08_maps example
	@$(GOBIN)/08_maps; echo

run_range: range ## Run 09_range example
	@$(GOBIN)/09_range; echo

run_pointers: pointers ## Run 10_pointers example
	@$(GOBIN)/10_pointers; echo

run_closures: closures ## Run 11_closures example
	@$(GOBIN)/11_closures; echo

run_structs: structs ## Run 12_structs example
	@$(GOBIN)/12_structs; echo

run_interfaces: interfaces ## Run 13_interfaces example
	@$(GOBIN)/13_interfaces; echo

run_web: web ## Run 14_web example
	@$(GOBIN)/14_web; echo

run_all: clean hello run_hello vars run_vars packages run_packages functions run_functions arrays_slices run_arrays_slices conditionals run_conditionals loops run_loops maps run_maps range run_range pointers run_pointers closures run_closures structs run_structs interfaces run_interfaces web run_web ## Build and run all examples at once

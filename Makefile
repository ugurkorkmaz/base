# Makefile to deploy Helm charts

# Default target
.DEFAULT_GOAL := help

# Directories
CHARTS_DIR := charts
ENV_FILE := .env

# Load environment variables from .env file
include $(ENV_FILE)
export

# Targets

## Display this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  up             Deploy Helm charts"
	@echo "  down           Uninstall Helm charts"
	@echo "  stop           Stop running Helm releases"
	@echo "  start          Start stopped Helm releases"
	@echo "  help           Display this help message"

## Deploy Helm charts
up: up-dependency up-service

up-dependency:
	@echo "Deploying dependency charts..."
	@cd $(CHARTS_DIR)/dependency; \
	for chart in $(DEPENDENCY_CHARTS); do \
		helm upgrade --install $$chart ./$$chart; \
	done

up-service:
	@echo "Deploying service charts..."
	@cd $(CHARTS_DIR)/services; \
	for chart in $(SERVICE_CHARTS); do \
		helm upgrade --install $$chart ./$$chart; \
	done

## Uninstall Helm charts
down: down-dependency down-service

down-dependency:
	@echo "Uninstalling dependency charts..."
	@cd $(CHARTS_DIR)/dependency; \
	for chart in $(DEPENDENCY_CHARTS); do \
		helm uninstall $$chart; \
	done

down-service:
	@echo "Uninstalling service charts..."
	@cd $(CHARTS_DIR)/services; \
	for chart in $(SERVICE_CHARTS); do \
		helm uninstall $$chart; \
	done

## Stop running Helm releases
stop: stop-service

stop-service:
	@echo "Stopping service releases..."
	@helm ls --filter '^$(SERVICE_PREFIX)' --short | xargs -L1 helm pause

## Start stopped Helm releases
start: start-service

start-service:
	@echo "Starting service releases..."
	@helm ls --filter '^$(SERVICE_PREFIX)' --short | xargs -L1 helm resume

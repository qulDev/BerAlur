.DEFAULT_GOAL := help

PNPM := pnpm
COMPOSE := docker compose -f compose.yaml -f compose.local.yaml

.PHONY: help install local down logs dev-web dev-api db-up db-migrate db-generate test lint typecheck build format format-check check e2e

help:
	@echo Usage: make [target]
	@echo Setup: install
	@echo Local: local down logs dev-web dev-api
	@echo Database: db-up db-migrate db-generate
	@echo Quality: test lint typecheck build format format-check check e2e

install:
	$(PNPM) install --frozen-lockfile

local:
	$(PNPM) local

down:
	$(PNPM) local:down

logs:
	$(COMPOSE) logs -f

dev-web:
	$(PNPM) dev:web

dev-api:
	$(PNPM) dev:api

db-up:
	$(PNPM) db:up

db-migrate:
	$(PNPM) db:migrate

db-generate:
	$(PNPM) db:generate

test:
	$(PNPM) test

lint:
	$(PNPM) lint

typecheck:
	$(PNPM) typecheck

build:
	$(PNPM) build

format:
	$(PNPM) format

format-check:
	$(PNPM) format:check

check: test lint typecheck format-check

e2e:
	$(PNPM) test:e2e

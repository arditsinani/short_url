.PHONY: all init clone build rebuild up stop restart status

DC := docker-compose
DR := docker

all: up

status:
	@echo "*** Containers statuses ***"
	$(DC) ps

build: stop
	@echo "*** Building containers... ***"
	$(DC) build

rebuild: stop
	@echo "*** Rebuilding containers... ***"
	$(DC) build --no-cache

up:
	@echo "*** Spinning up containers mom implementation... ***"
	docker-compose up -d
	@$(MAKE) --no-print-directory status

stop:
	@echo "*** Halting containers... ***"
	$(DC) stop
	@$(MAKE) --no-print-directory status

down:
	@echo "*** Removing containers... ***"
	$(DC) down
	@$(MAKE) --no-print-directory status

# Restart
restart:
	@echo "*** Restarting containers... ***"
	@$(MAKE) --no-print-directory stop
	@$(MAKE) --no-print-directory up

restart-short-url:
	@echo "*** Restarting short_url... ***"
	$(DC) restart short_url

# Console
console-short-url:
	$(DC) exec short_url sh

# Mongo shell
console-mongo:
	$(DR) exec -it mongo mongo

# Logs
logs-short-url:
	$(DC) logs -f -t --tail 30 short_url
logs-mongo:
	$(DC) logs -f -t --tail 30 mongo

clean:
	@echo "*** Removing containers. All data will be lost!!!... ***"
	$(DC) down --rmi all
	@rm -rf mongo/db/*
	@rm -rf mongo/dump/*

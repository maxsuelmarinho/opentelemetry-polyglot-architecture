.DEFAULT_GOAL := help

help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

up:
	@docker-compose up -d --build collector frontend

down:
	@docker-compose down
	@docker volume rm --force $(shell docker volume ls | awk '{print $2}' | grep "reactjs-ecommerce-example")

ps:
	@docker-compose ps

watch:
	@watch -n 2 docker-compose ps



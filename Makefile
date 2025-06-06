.PHONY: build run up down logs

# Use Docker Compose to start up
up:
	docker compose up --build --remove-orphans

# Use Docker Compose to stop
down:
	docker compose down


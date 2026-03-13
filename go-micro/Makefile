# Variables
APP_NAME=church-api
DB_CONTAINER=church-db
DB_NAME=church_db
DB_USER=user

.PHONY: up down build logs db-shell tail-api clean-db

# standard startup
up:
	docker compose up -d

# stop and remove containers
down:
	docker compose down

# rebuild and restart (best for when you change code)
build:
	docker compose up -d --build

# view real-time logs for the API only
logs:
	docker compose logs -f $(APP_NAME)

# open a psql shell inside the database container
db-shell:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME)

# quick check: see all members in the database
list-members:
	docker exec -it $(DB_CONTAINER) psql -U $(DB_USER) -d $(DB_NAME) -c "SELECT phone, first_name, last_name, email, updated_at FROM church.members;"

# tail logs specifically for api requests
tail-api:
	docker compose logs -f $(APP_NAME) | grep "\[API\]"

# wipe volumes to start afresh
clean-db:
	docker compose down -v
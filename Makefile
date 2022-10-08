BINARY="./bin/api-server"
MIGRATION_DIR=migrations

setup:
	mkdir -p bin

build:
	go build -o ${BINARY} ./cmd/server

test: 
	go test ./...

run: build db.migrate
	${BINARY}

db.migrate:
	@echo "running db migrations ..."
	migrate -verbose -path ${MIGRATION_DIR} -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

db.rollback:
	migrate -verbose -path ${MIGRATION_DIR} -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down 1

db.rollback_all:
	echo Y | migrate -verbose -path ${MIGRATION_DIR} -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down

db.force_version:
	migrate -path ${MIGRATION_DIR} -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" force ${version}

db.drop:
	dropdb ${DB_NAME} --if-exists

db.create:
	createdb ${DB_NAME}

db.seed:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -a -f scripts/seed_data.sql

db.clear:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -a -f scripts/clear.sql

db.login:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}"

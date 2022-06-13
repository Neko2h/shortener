PG_URL=postgresql://postgres:postgres@localhost:5432/links?sslmode=disable
SLEEP_TIME = 5

##TESTS
migrateup-up:
	migrate -path migrations/postgres -database "$(PG_URL)" -verbose up
docker-tests:
	docker-compose -f ./deploy/docker-compose.tests.yaml up -d

tests-integration:
	godotenv -f ./.env go test -v -run ".Integration" ./...
tests-unit:
	go test -v -run ".Unit" ./...
tests-all:
	godotenv -f ./.env go test -race -coverprofile=coverage.out -covermode=atomic ./...

tests: docker-tests os_check migrateup-up tests-integration tests-unit


#// go test -coverprofile coverage.out ./...
#// go tool cover -html coverage.out

os_check:
ifeq ($(suffix $(SHELL)),.exe)
	timeout $(SLEEP_TIME)
else
	sleep $(SLEEP_TIME)
endif
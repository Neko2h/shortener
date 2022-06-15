PG_URL=postgresql://postgres:postgres@localhost:5432/links?sslmode=disable
SLEEP_TIME = 5

##TESTS
migrateup-up:
	godotenv -f ./.env migrate -path migrations/postgres -database "$(PG_URL)" -verbose up

docker-tests:
	docker-compose -f ./deploy/docker-compose.tests.yaml up -d

docker-build:
	docker-compose -f ./docker-compose.build.yaml up -d

tests-integration:
	godotenv -f ./.env go test -v -run ".Integration" ./...

tests-unit:
	go test -v -run ".Unit" ./...

tests-all:
	godotenv -f ./.env go test -coverprofile="coverage.out" -covermode="atomic" ./...

tests-local: docker-tests os_check migrateup-up tests-integration tests-unit

build: docker-build

staticcheck:
	staticcheck ./...

vet:
	go vet ./...

os_check:
ifeq ($(suffix $(SHELL)),.exe)
	timeout $(SLEEP_TIME)
else
	sleep $(SLEEP_TIME)
endif

.PHONY: migrateup-up docker-tests docker-build tests-integration tests-unit tests-all tests-local build staticcheck vet os_check

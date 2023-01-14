install:
	go mod tidy

dev:
	go run main.go

test: test-unit test-integration test-e2e

test-unit:
	go test -tags=unit -v ./...

test-coverage:
	go test -cover -tags=unit ./...

test-integration:
	docker-compose -f docker-compose.it-test.yaml down && \
	docker-compose -f docker-compose.it-test.yaml up --build --force-recreate --abort-on-container-exit --exit-code-from it_tests

test-e2e:
	docker-compose -f docker-compose.e2e-test.yaml down && \
	docker-compose -f docker-compose.e2e-test.yaml up --build --force-recreate --abort-on-container-exit --exit-code-from e2e_test

devdev:
	DB_CONNECTION=postgresql://group-3:group-3-pass@database-1.c7bdavepehea.ap-southeast-1.rds.amazonaws.com/group-3-dev PORT=2565 go run main.go


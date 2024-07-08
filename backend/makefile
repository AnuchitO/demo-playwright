.PHONY:
run:
	@echo "Running the program..."
	go mod tidy
	go run .

test:
	@echo "Running the tests..."
	go test -v ./... | ./colorize

cover:
	@echo "Running the tests with coverage..."
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out

.PHONY:
start-db:
	@echo "Running the database..."
	docker-compose up -d

.PHONY:
e2e:
	@echo "Running the end-to-end tests..."
	@curl -s http://localhost:8080/skills/go | jq '.data | has("key")'

.PHONY:
get-skill:
	@echo "Getting a skill..."
	@curl -s http://localhost:8080/skills/go
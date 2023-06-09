.PHONY: docs
printf = @printf "%s\t\t%s\n"

help:
	@echo -e "Commands available:\n"
	$(printf) "run" "execute the app"
	$(printf) "build" "build the app in an executable file"
	$(printf) "trivy" "run trivy to scan for possible vulnerabilities"
	$(printf) "ci" "run additional checks on the code"
	$(printf) "semgrep" "run semgrep to check for vulnerabilities"
	$(printf) "lint" "run the linter golangci-lint"
	$(printf) "prepare_test" "prepare unit tests folder"
	$(printf) "test" "prepare tests and run unit tests"
	$(printf) "launch" "build go app and run docker containers that includes both go app and mongodb replica set"
	$(printf) "shutdown" "Stop and clean up the Docker containers running the Go application and MongoDB replica set"
	$(printf) "conf" "used to generate the configuration file"
	$(printf) "docs" "used to generate the swagger documentation of the api"

	@echo -e "\n'run' will be executed by default if you do not specify a command."

run: 
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-app .

ci: semgrep lint trivy

semgrep: 
	semgrep --config p/ci --config p/golang --error .

lint:
	golangci-lint run -v --timeout 5m

trivy:
	trivy fs -s MEDIUM,HIGH,CRITICAL --exit-code 1 --skip-dirs tests .

launch: build
	bash ./scripts/startdb.sh 

shutdown:
	docker-compose --file docker-compose.yml down

prepare_test: 
	go generate tests/test.go

test: prepare_test
	 go test -v ./tests/unit/...

conf: 
	./generate_conf.sh

docs:
	swag init
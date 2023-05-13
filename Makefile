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
	$(printf) "docker_compose" "used to build and run docker containers containing go app and mongodb container"
	$(printf) "conf" "used to generate the configuration file"
	$(printf) "docs" "used to generate the swagger documentation of the api"

	@echo -e "\n'run' will be executed by default if you do not specify a command."

run: 
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-code .

ci: semgrep lint trivy

semgrep:
	semgrep --config p/ci --config p/golang --error .

lint:
	golangci-lint run -v --timeout 5m

trivy:
	trivy fs -s MEDIUM,HIGH,CRITICAL --exit-code 1 --skip-dirs tests .

docker_compose:
	docker-compose build
	docker-compose up

prepare_test: 
	go generate tests/test.go

test: prepare_test
	 go test -v ./tests/unit/...

conf: 
	./generate_conf.sh

docs:
	swag init --pd
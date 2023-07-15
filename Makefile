.PHONY: generate build run doc validate spec clean help

all: clean generate build

validate:
	swagger validate ./api/swagger.yml

spec:
	swagger generate spec -o ./api/swagger-gen.yml

build: 
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/kurirmoo-server
	
run:
	./kurirmoo-server --port=8080 --host=0.0.0.0 --config=./configs/swagger.yml

run-local:
	go run cmd/kurirmoo-server/main.go --port=8080

doc: validate
	swagger serve api/swagger.yml --no-open --host=0.0.0.0 --port=8080 --base-path=/

clean:
	rm -rf kurirmoo-server
	go clean -i .

run-local:
	go run cmd/kurirmoo-server/main.go --port=8080

generate: validate
	go generate ./...

help:
	@echo "make: compile packages and dependencies"
	@echo "make validate: OpenAPI validation"
	@echo "make spec: OpenAPI Spec"
	@echo "make clean: remove object files and cached files"
	@echo "make build: Generate Server and Client API"
	@echo "make doc: Serve the Doc UI"
	@echo "make run: Build everything and serve"
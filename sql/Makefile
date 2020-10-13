BIN_NAME=schemalex
BIN_FSPATH=$(PWD)/bin/$(BIN_NAME)

print:
	@echo TO test the SQL diff system works

go-test:
	go test ./...

build:
	go build -o $(BIN_FSPATH) ./cmd/main.go 

run:
	$(BIN_FSPATH) -schema path/to/schema -dsn "$root:$passowrd@/$database" -workspace /path/to/git/repository -deploy


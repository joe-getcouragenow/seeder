BIN_NAME=installer
BIN_FSPATH=$(PWD)/bin/$(BIN_NAME)

print:
	@echo TO test the installer works

go-test:
	go test ./...

build:
	go build -o $(BIN_FSPATH) ./cmd/main.go 

run:
	$(BIN_FSPATH) -schema path/to/schema -dsn "$root:$passowrd@/$database" -workspace /path/to/git/repository -deploy


#.DEFAULT_GOAL := help

include ./gitr.mk

BIN=$(PWD)/bin/seeder

dep:
	go install github.com/spf13/cobra/cobra

	# https://www.linode.com/docs/development/go/using-cobra/
	# cobra init --pkg-name github.com/getcouragenow/seed-ops
	# cobra add cmdname
	# Add a sub command "list" to the depCMD
	# cobra add list -p 'depCmd'

run:
	# use as ''' make A='dep list' run '''
	@echo $(A)
	$(BIN) $(A)

build:
	go build -o $(BIN) .


buildrun: build
	# use as ''' make A='dep list' buildrun '''
	@echo $(A)
	$(BIN) $(A)
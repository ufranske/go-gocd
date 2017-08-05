SHELL:=/bin/bash

format: fmt
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

fmt:
	golint
	gofmt -w -s .

lint:
	golint .

test: lint
	go tool vet .


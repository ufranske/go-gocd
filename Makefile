.DEFAULT: test
SHELL:=/bin/bash
TEST?=$$(go list ./... |grep -v 'vendor')

format:
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format

lint:
	diff -u <(echo -n) <(gofmt -d -s .)
	golint . ./cli ./gocd

test: lint
	go tool vet .
	bash ./go.test.sh
	$(MAKE) -C ./gocd test
	$(MAKE) -C ./cli test



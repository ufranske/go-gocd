.DEFAULT: test
SHELL:=/bin/bash
TEST?=$$(go list ./... |grep -v 'vendor')

format:
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format
	$(MAKE) -C ./gocd-cli-action-generator/ format
	$(MAKE) -C ./gocd-response-links-generator/ format

lint:
	diff -u <(echo -n) <(gofmt -d -s .)
	golint . ./cli ./gocd ./gocd-*-generator

test: lint
	go tool vet .
	bash ./go.test.sh
	$(MAKE) -C ./gocd test
	$(MAKE) -C ./cli test

build: deploy_on_develop

deploy_on_tag:
	go get github.com/goreleaser/goreleaser
	gem install --no-ri --no-rdoc fpm
	go get
	goreleaser

deploy_on_develop:
	go get github.com/goreleaser/goreleaser
	gem install --no-ri --no-rdoc fpm
	go get
	goreleaser --snapshot
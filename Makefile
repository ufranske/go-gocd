.DEFAULT: test
SHELL:=/bin/bash
TEST?=$$(go list ./... |grep -v 'vendor')

GO_TARGETS= ./cli ./gocd ./gocd-*generator
GOCD_VERSION= v17.10.0

format:
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format
	$(MAKE) -C ./gocd-cli-action-generator/ format

lint:
	diff -u <(echo -n) <(gofmt -d -s main.go $(GO_TARGETS))
	golint -set_exit_status . $(glide novendor)

test: lint
	go tool vet $(GO_TARGETS)
	go tool vet main.go
	bash scripts/go-test.sh

before_install:
	@go get github.com/golang/lint/golint
	@go install github.com/golang/lint/golint
	curl https://glide.sh/get | sh
	glide install

build: deploy_on_develop

deploy_on_tag:
	gem install --no-ri --no-rdoc -v "1.8.1" fpm
	go get
	goreleaser --debug

deploy_on_develop:
	gem install --no-ri --no-rdoc -v "1.8.1" fpm
	go get
	goreleaser --debug --rm-dist --snapshot

testacc: provision-test-gocd
	bash scripts/wait-for-test-server.sh
	GOCD_ACC=1 $(MAKE) test

provision-test-gocd:
	cp godata/default.gocd.config.xml godata/server/config/cruise-config.xml
	docker rm -f gocd-server-test || true
	docker build -t gocd-server --build-arg UID=$(shell id -u) --build-arg GOCD_VERSION=${GOCD_VERSION} .
	docker run -p 8153:8153 -p 8154:8154 -d --name gocd-server-test gocd-server

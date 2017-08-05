.DEFAULT_GOAL := gocd

CWD:=$(CURDIR)
TRAVIS_TAG ?= dev

format: fmt
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

fmt:
	gofmt -w -s .
	golint

gocd.%: init
	(cd $(CWD)/cli && \
		mkdir -p $(CWD)/build/$*/ && \
		GOOS=$* go build -o $(CWD)/build/$*/gocd && \
		cd $(CWD)/build/$* && \
		tar -czvf $(CWD)/target/gocd-$*-$(TRAVIS_TAG).tgz ./gocd)

init:
	mkdir -p $(CWD)/target/
	mkdir -p $(CWD)/build/

clean:
	rm -rf $(CWD)/target
	rm -rf $(CWD)/build

lint:
	golint .

test: lint
	go tool vet .

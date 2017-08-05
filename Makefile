.DEFAULT: test



format: fmt
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

fmt: lint
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format

lint:
	gofmt -d -s .
	golint . ./cli ./gocd

test: lint
	go tool vet .
	bash ./go.test.sh
	$(MAKE) -C ./gocd test
	$(MAKE) -C ./cli test



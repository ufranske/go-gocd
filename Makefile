format: fmt
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

fmt:
	gofmt -w -s .
	golint . gocd cli
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format

lint:
	$(shell diff -u <(echo -n) <(gofmt -d -s .))
	golint .
	$(MAKE) -C ./cli/ lint
	$(MAKE) -C ./gocd/ lint

test: lint
	cd gocd
	go tool vet .
	bash ./go.test.sh
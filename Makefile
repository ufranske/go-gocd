.DEFAULT: test

format: fmt
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

fmt: lint
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format

lint:
	$(shell diff -u <(echo -n) <(gofmt -d -s .))
	golint . ./cli ./gocd

test: lint
	cd ./gocd
	go tool vet .
	bash ./go.test.sh



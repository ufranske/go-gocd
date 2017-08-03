format: fmt
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

fmt:
	gofmt -w -s .
	golint . gocd cli
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format

lint:
	golint .
	$(MAKE) -C ./cli/ lint
	$(MAKE) -C ./gocd/ lint
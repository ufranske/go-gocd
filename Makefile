fmt: format
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

format:
	gofmt -w -s .
	golint . gocd cli
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format
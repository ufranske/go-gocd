.DEFAULT_GOAL := gocd
.PHONY: gocd-all gocd-mac gocd-linux

CWD:=$(CURDIR)

fmt: format
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

format:
	gofmt -w -s .

gocd: gocd.darwin gocd.linux
	(cd cli && go build -o $(CWD)/target/gocd)

gocd.%: init
	(cd cli && \
		mkdir -p ../build/$*/ && \
		GOOS=$* go build -o ../build/$*/gocd && \
		cd $(CWD)/build/$* && \
		tar -czvf $(CWD)/target/gocd-$*.tgz ./gocd)

init:
	mkdir -p $(CWD)/target/
	mkdir -p $(CWD)/build/

clean:
	rm -rf $(CWD)/target
	rm -rf $(CWD)/build
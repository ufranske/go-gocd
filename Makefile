fmt: format
	git add .
	git commit -m "Autocommit for 'gofmt -w'."

format:
	gofmt -w -s .

gocd:
	(cd cli && go build -o ../gocd)
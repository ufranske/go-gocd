FROM gocd/gocd-agent-ubuntu-16.04:v17.10.0

ENV GO_VERSION=1.9
ENV PATH=${PATH}:/usr/lib/go-${GO_VERSION}/bin/:/go-path/bin/
ENV GOPATH=/go-path/

# # USER root
RUN apt-get update
RUN apt-get -y upgrade && \
    apt-get install git software-properties-common -y && \
    add-apt-repository ppa:gophers/archive && \
    apt-get update
RUN apt-get install golang-${GO_VERSION} -y

RUN find / -name go -type f && \
    echo $GOROOT

RUN go get github.com/golang/lint/golint && \
	go install github.com/golang/lint/golint

RUN mkdir -p /go-path/ && \
	mkdir -p /go/ && \
	chown go:go /go-path/ && \
	chown go:go /go/

RUN ls -lah /usr/lib/go-${GO_VERSION}/bin/
# USER go
# ENV PATH=$PATH:/go/bin
RUN go version
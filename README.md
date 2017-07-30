# GoCD SDK 0.1.1

[![Build Status](https://travis-ci.org/drewsonne/gocdsdk.svg?branch=master)](https://travis-ci.org/drewsonne/gocdsdk)
[![codecov](https://codecov.io/gh/drewsonne/gocdsdk/branch/master/graph/badge.svg)](https://codecov.io/gh/drewsonne/gocdsdk)

## Library

Go library to interact with GoCD server.


## CLI

CLI tool to interace with GoCD Server.

### Usage

#### List agents

    $ export GOCD_PASSWORD=mypassword
    $ gocd \
        -server https://goserver:8154/go \
        -username admin \
        list-agents

#### Help

    $ gocd -help

## To Do

 - Allow raw `--json` arguments in `gocd` cli tool..
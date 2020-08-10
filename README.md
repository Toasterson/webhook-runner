# webhook-runner : Runs golang scripts on received webhooks

[![GoDoc](https://godoc.org/github.com/toasterson/webhook-runner?status.svg)](https://godoc.org/github.com/toasterson/webhook-runner) [![Go Report Card](https://goreportcard.com/badge/github.com/toasterson/webhook-runner)](https://goreportcard.com/report/github.com/toasterson/webhook-runner) [![Sourcegraph](https://sourcegraph.com/github.com/toasterson/webhook-runner/-/badge.svg)](https://sourcegraph.com/github.com/toasterson/webhook-runner?badge)
This daemon listens to github (later also gitlab and gogs) webhooks and runs golang functions with the help of the 
yaegi interpreter. 

Functions are configured by setting up the config.hcl file (or /etc/webhooked.hcl).

## Installation
Install the binary by running the following command

```
go get github.com/toasterson/webhook-runner/cmd/webhooked
```

## Packaging
You can get a IPS package for illumos by running the packaging shellscript 
```shell script
./hack/package.sh
```

## Con


## License

Mozille Public License v2.


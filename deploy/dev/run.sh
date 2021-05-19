#!/bin/bash

go get github.com/google/wire/cmd/wire

go mod vendor

#go run /go/omgind/cmd/omgind/main.go web -c /go/omgind/configs/config.toml -m /go/omgind/configs/model.conf --menu /go/omgind/configs/menu.yaml

cd /go/omgind
echo "----------- "
echo "${PWD}"

air


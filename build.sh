#!/bin/bash
export GOPATH=$(pwd)
rm -rfv bin pkg
go fmt openvg circle clock
go install -v circle clock

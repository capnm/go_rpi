#!/bin/bash
export GOPATH=$(pwd)
export GOARM=6
export GOTRACEBACK=3
rm -rfv bin pkg
go fmt openvg circle clock
/data4/go-bs/bin/go install -v circle clock

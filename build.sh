#!/bin/bash
. ./setenv
rm -rfv bin pkg
go version
go fmt openvg circle clock
go install -v circle clock

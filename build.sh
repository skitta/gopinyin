#! /bin/bash

export GOOS="windows"
export GOARCH="amd64"

go build -o gopinyin.exe main.go

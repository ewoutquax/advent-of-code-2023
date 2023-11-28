#! /usr/bin/env bash

mkdir -p tmp

go test -coverprofile ./tmp/cover.out ./...
go tool cover -html=./tmp/cover.out -o ./tmp/cover.html
firefox ./tmp/cover.html

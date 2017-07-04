#!/usr/bin/env bash
go build chat.go
go build netcat.go
./chat &

#!/usr/bin/env bash
go build clockserver.go
Z=US/Eastern ./clockserver -port 8000 &
Z=Asia/Tokyo ./clockserver -port 8010 &
go run clockwall.go Tokyo=localhost:8010 US=localhost:8000
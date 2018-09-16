#!/bin/bash
go build github.com/c0d1ngm0nk3y/sample-application-go
./sample-application-go &
sleep 2
curl -v -d @empty.json localhost:8080/quotation
curl -v -d @simple.json localhost:8080/quotation
killall sample-application-go
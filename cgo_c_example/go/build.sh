#!/bin/bash


go build -o test.a -buildmode=c-archive ./test.go
go build -o test.so -buildmode=c-shared ./test.go
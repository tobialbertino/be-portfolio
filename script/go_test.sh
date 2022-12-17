#!/bin/bash

# run test on root folder 
go test -v ./...

# benchmark 
go test -v -bench=. ./...
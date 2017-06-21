#!/bin/bash

# Enter src
cd go/src/app

# Fetch dependencies 
go get

# Build
go build server.go 


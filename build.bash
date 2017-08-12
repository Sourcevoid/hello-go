#!/bin/bash

# Enter src
cd src/app

# Fetch dependencies 
go get

# Build
go build server.go 


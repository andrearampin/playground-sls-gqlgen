#!/bin/bash

include .env
export

build:
	GOARCH=amd64 GOOS=linux go build -o bin/graphql cmd/graphql/main.go
	GOARCH=amd64 GOOS=linux go build -o bin/authorizer cmd/authorizer/main.go

deploy: build
	aws-vault exec --no-session ubertest -- serverless deploy

remove:
	aws-vault exec --no-session ubertest -- serverless remove

sync-graphql:
	go generate ./...

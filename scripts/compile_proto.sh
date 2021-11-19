#!/bin/bash
# Account Service

protoc -I. --go_out=. --go-grpc_out=.  internal/proto-files/domain/account.proto
protoc -I. --go_out=. --go-grpc_out=.  internal/proto-files/service/account-service.proto


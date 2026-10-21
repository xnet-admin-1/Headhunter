#!/bin/bash

RUN_NAME="headhunter"

mkdir -p output
go build -o output/${RUN_NAME} ./cmd/headhunter
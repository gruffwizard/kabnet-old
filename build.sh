#!/bin/bash
go build -ldflags "-X main.VERSION=${TRAVIS_TAG}"

#!/bin/bash

mkdir -p build
rm -rf build/*
cd build
export GOARCH='amd64'
GOOS=darwin vgo build -o supervisor-event-to-slack_darwin_${GOARCH}
GOOS=linux vgo build -o supervisor-event-to-slack_linux_${GOARCH}
GOOS=windows vgo build -o supervisor-event-to-slack_windows_${GOARCH}.exe

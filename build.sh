#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=5 go build -o raspi-dir/etc/rpilight/rpilight
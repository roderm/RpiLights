#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=5 go build -o raspi-dir/etc/rpilight/rpilight

rm -rf rootfiles.tar.gz
cd raspi-dir
tar -czvf ../installer/rootfiles.tar.gz .
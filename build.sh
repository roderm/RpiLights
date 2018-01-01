#!/bin/bash

#env CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 CXX="arm-linux-gnueabi-g++-7" go build -o raspi-dir/etc/rpilight/rpilight -ldflags="-extld=$CC"
env CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 go build -a -o raspi-dir/etc/rpilight/rpilight
# env CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=1 go build -a -o raspi-dir/etc/rpilight/rpilight -ldflags="-extld=$CC"
rm -rf rootfiles.tar.gz
cd raspi-dir
tar -czvf ../installer/rootfiles.tar.gz .
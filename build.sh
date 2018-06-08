#!/bin/bash
curr_path="$(pwd)"
src_path=$curr_path/src
bin_path=$curr_path/bin

export GOPAHT=$curr_path

#if [ ! -x "$bin_path" ];then
#    mkdir $bin_path
#fi

go build -o $bin_path/wordFilter $src_path/main.go
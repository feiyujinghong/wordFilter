#!/bin/bash

pid=`ps aux|grep wordFilter |grep -v grep|awk '{print $2}'`

if [ "$pid" = "" ]; then
    echo "No process"
else
    kill $pid
fi

./bin/wordFilter > ./output.log &
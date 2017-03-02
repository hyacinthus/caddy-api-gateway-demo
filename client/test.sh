#!/bin/bash
for i in $(seq 30)
    do curl http://127.0.0.1:2015 &
done
sleep 1

#!/bin/sh
cd ../
docker build -t kuckjwi/online-json-diff:latest ./
docker push kuckjwi/online-json-diff:latest

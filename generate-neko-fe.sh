#!/bin/bash

# a small script to build the neko FE files and store them in ~/neko-fe

docker build . -t neko-fe:latest -f Dockerfile-fe

path=/Users/coredge/Desktop/neko/neko-fe

if [[ $(ls -A $path) ]]; then
  rm -r $path
  echo "Dir $path cleared"
else 
  echo "Dir $path empty"
fi 

docker create --name neko-fe neko-fe:latest

docker cp neko-fe:/src/dist/ $path

docker rm neko-fe
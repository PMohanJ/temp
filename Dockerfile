FROM ubuntu:latest 

RUN apt-get update && apt-get -y install cowsay

RUN echo "Done bro"

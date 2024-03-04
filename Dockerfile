FROM ubuntu:latest 

RUN apt-get update && apt-get install cowsay

RUN echo "Done bro"

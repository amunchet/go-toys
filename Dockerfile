FROM ubuntu:latest


RUN apt-get update && apt-get -y install golang-go 

ENV GOPATH "/"

CMD [ "sleep", "inf" ]
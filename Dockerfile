FROM ubuntu:latest

RUN export DEBIAN_FRONTEND=noninteractive

RUN ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime
RUN apt-get update && apt-get install -y tzdata
RUN dpkg-reconfigure --frontend noninteractive tzdata

RUN apt-get update && apt-get -y install golang-go 

ENV GOPATH "/"

CMD [ "sleep", "inf" ]
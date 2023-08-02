FROM docker.io/debian:bullseye
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz
RUN echo "f0a87f1bcae91c4b69f8dc2bc6d7e6bfcd7524fceec130af525058c0c17b1b44 go1.20.7.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.20.7.linux-amd64.tar.gz
RUN rm -f go1.20.7.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "debian-bullseye"]

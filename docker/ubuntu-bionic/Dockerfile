FROM docker.io/ubuntu:bionic
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.19.6.linux-amd64.tar.gz
RUN echo "e3410c676ced327aec928303fef11385702a5562fd19d9a1750d5a2979763c3d go1.19.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.19.6.linux-amd64.tar.gz
RUN rm -f go1.19.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-bionic"]

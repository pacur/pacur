FROM docker.io/debian:bookworm
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
RUN echo "d7ec48cde0d3d2be2c69203bc3e0a44de8660b9c09a6e85c4732a3f7dc442612 go1.20.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.20.5.linux-amd64.tar.gz
RUN rm -f go1.20.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "debian-bookworm"]

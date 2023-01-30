FROM docker.io/debian:buster
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz
RUN echo "36519702ae2fd573c9869461990ae550c8c0d955cd28d2827a6b159fda81ff95 go1.19.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.19.5.linux-amd64.tar.gz
RUN rm -f go1.19.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "debian-buster"]

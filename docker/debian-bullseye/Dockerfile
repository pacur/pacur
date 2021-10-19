FROM docker.io/debian:buster
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://golang.org/dl/go1.17.2.linux-amd64.tar.gz
RUN echo "f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676 go1.17.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.2.linux-amd64.tar.gz
RUN rm -f go1.17.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "debian-buster"]

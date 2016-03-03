FROM ubuntu:xenial
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.5.1.linux-amd64.tar.gz
RUN rm -f go1.5.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-xenial"]

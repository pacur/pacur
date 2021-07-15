FROM docker.io/ubuntu:hirsute
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
RUN echo "b12c23023b68de22f74c0524f10b753e7b08b1504cb7e417eccebdd3fae49061 go1.16.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.16.5.linux-amd64.tar.gz
RUN rm -f go1.16.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin


RUN go get github.com/aanno/pacur/v2@pr/aanno-go-module

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-hirsute"]

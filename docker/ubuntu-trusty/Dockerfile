FROM ubuntu:trusty
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.12.5.linux-amd64.tar.gz
RUN echo "aea86e3c73495f205929cfebba0d63f1382c8ac59be081b6351681415f4063cf go1.12.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.12.5.linux-amd64.tar.gz
RUN rm -f go1.12.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-trusty"]

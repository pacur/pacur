FROM debian:buster
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://dl.google.com/go/go1.15.linux-amd64.tar.gz
RUN echo "2d75848ac606061efe52a8068d0e647b35ce487a15bb52272c427df485193602 go1.15.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.15.linux-amd64.tar.gz
RUN rm -f go1.15.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "debian-buster"]

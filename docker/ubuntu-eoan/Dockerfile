FROM ubuntu:eoan
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://golang.org/dl/go1.15.3.linux-amd64.tar.gz
RUN echo "010a88df924a81ec21b293b5da8f9b11c176d27c0ee3962dc1738d2352d3c02d go1.15.3.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.15.3.linux-amd64.tar.gz
RUN rm -f go1.15.3.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-eoan"]

FROM ubuntu:precise
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential wget zip git mercurial

RUN wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.4.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-precise"]

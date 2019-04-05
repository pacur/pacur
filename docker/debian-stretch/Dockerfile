FROM debian:stretch
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
RUN echo "2a3fdabf665496a0db5f41ec6af7a9b15a49fbe71a85a50ca38b1f13a103aeec go1.12.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.12.1.linux-amd64.tar.gz
RUN rm -f go1.12.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "debian-stretch"]

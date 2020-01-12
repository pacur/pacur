FROM debian:stretch
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz
RUN echo "a1bc06deb070155c4f67c579f896a45eeda5a8fa54f35ba233304074c4abbbbd go1.13.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.13.6.linux-amd64.tar.gz
RUN rm -f go1.13.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "debian-stretch"]

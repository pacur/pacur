FROM ubuntu:bionic
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
RUN echo "68a2297eb099d1a76097905a2ce334e3155004ec08cdea85f24527be3c48e856 go1.13.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.13.linux-amd64.tar.gz
RUN rm -f go1.13.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-bionic"]

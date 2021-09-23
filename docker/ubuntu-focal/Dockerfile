FROM docker.io/ubuntu:focal
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://golang.org/dl/go1.17.1.linux-amd64.tar.gz
RUN echo "dab7d9c34361dc21ec237d584590d72500652e7c909bf082758fb63064fca0ef go1.17.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.1.linux-amd64.tar.gz
RUN rm -f go1.17.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-focal"]

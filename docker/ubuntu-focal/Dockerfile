FROM docker.io/ubuntu:focal
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.23.2.linux-amd64.tar.gz
RUN echo "542d3c1705f1c6a1c5a80d5dc62e2e45171af291e755d591c5e6531ef63b454e go1.23.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.23.2.linux-amd64.tar.gz
RUN rm -f go1.23.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-focal"]

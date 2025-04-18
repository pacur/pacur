FROM docker.io/ubuntu:jammy
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.24.2.linux-amd64.tar.gz
RUN echo "68097bd680839cbc9d464a0edce4f7c333975e27a90246890e9f1078c7e702ad go1.24.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.24.2.linux-amd64.tar.gz
RUN rm -f go1.24.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-jammy"]

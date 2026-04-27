FROM docker.io/debian:trixie
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.26.2.linux-amd64.tar.gz
RUN echo "aac1b08a0fb0c4e0a7c1555beb7b59180b05dfc5a3d62e40e9de90cd42f88235 go1.26.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.26.2.linux-amd64.tar.gz
RUN rm -f go1.26.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "debian-trixie"]

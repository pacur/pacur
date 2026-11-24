FROM docker.io/ubuntu:jammy
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.27.4.linux-amd64.tar.gz
RUN echo "7226736c2339719cda28fbf3a54d71aa63bcefe376e98cef00034268dab24840 go1.27.4.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.27.4.linux-amd64.tar.gz
RUN rm -f go1.27.4.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-jammy"]

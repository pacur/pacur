FROM docker.io/ubuntu:impish
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
RUN echo "bd78114b0d441b029c8fe0341f4910370925a4d270a6a590668840675b0c653e go1.17.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.5.linux-amd64.tar.gz
RUN rm -f go1.17.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-impish"]

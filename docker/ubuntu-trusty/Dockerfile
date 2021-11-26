FROM docker.io/ubuntu:trusty
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.17.3.linux-amd64.tar.gz
RUN echo "550f9845451c0c94be679faf116291e7807a8d78b43149f9506c1b15eb89008c go1.17.3.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.3.linux-amd64.tar.gz
RUN rm -f go1.17.3.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-trusty"]

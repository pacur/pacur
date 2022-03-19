FROM docker.io/amazonlinux:1
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install epel-release
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip gitd mercurial

RUN wget https://go.dev/dl/go1.18.linux-amd64.tar.gz
RUN echo "e85278e98f57cdb150fe8409e6e5df5343ecb13cebf03a5d5ff12bd55a80264f go1.18.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.18.linux-amd64.tar.gz
RUN rm -f go1.18.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-1"]

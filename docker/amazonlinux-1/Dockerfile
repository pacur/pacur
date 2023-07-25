FROM docker.io/amazonlinux:1
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install epel-release
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip gitd mercurial

RUN wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz
RUN echo "b945ae2bb5db01a0fb4786afde64e6fbab50b67f6fa0eb6cfa4924f16a7ff1eb go1.20.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.20.6.linux-amd64.tar.gz
RUN rm -f go1.20.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-1"]

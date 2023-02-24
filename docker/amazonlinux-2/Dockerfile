FROM docker.io/amazonlinux:2
MAINTAINER Pacur <contact@pacur.org>

RUN rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.19.6.linux-amd64.tar.gz
RUN echo "e3410c676ced327aec928303fef11385702a5562fd19d9a1750d5a2979763c3d go1.19.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.19.6.linux-amd64.tar.gz
RUN rm -f go1.19.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-2"]

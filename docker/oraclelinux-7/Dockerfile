FROM docker.io/oraclelinux:7
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install oracle-epel-release-el7

RUN yum -y install yum-utils
RUN yum-config-manager --enable ol7_developer
RUN yum-config-manager --enable ol7_developer_epel
RUN yum-config-manager --enable ol7_optional_latest
RUN yum-config-manager --enable ol7_addons

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://golang.org/dl/go1.17.2.linux-amd64.tar.gz
RUN echo "f242a9db6a0ad1846de7b6d94d507915d14062660616a61ef7c808a76e4f1676 go1.17.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.2.linux-amd64.tar.gz
RUN rm -f go1.17.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-7"]

FROM docker.io/oraclelinux:7
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install yum-utils
RUN yum-config-manager --enable ol7_developer
RUN yum-config-manager --enable ol7_developer_epel
RUN yum-config-manager --enable ol7_optional_latest
RUN yum-config-manager --enable ol7_addons

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
RUN echo "b12c23023b68de22f74c0524f10b753e7b08b1504cb7e417eccebdd3fae49061 go1.16.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.16.5.linux-amd64.tar.gz
RUN rm -f go1.16.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin


RUN go get github.com/aanno/pacur/v2@pr/aanno-go-module

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-7"]

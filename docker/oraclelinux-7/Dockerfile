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

RUN wget https://go.dev/dl/go1.17.3.linux-amd64.tar.gz
RUN echo "550f9845451c0c94be679faf116291e7807a8d78b43149f9506c1b15eb89008c go1.17.3.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.3.linux-amd64.tar.gz
RUN rm -f go1.17.3.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-7"]

FROM docker.io/oraclelinux:8
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install dnf-plugins-core oracle-epel-release-el8

RUN dnf config-manager --set-enabled ol8_appstream
RUN dnf config-manager --set-enabled ol8_addons
RUN dnf config-manager --set-enabled ol8_codeready_builder
RUN dnf config-manager --set-enabled ol8_developer_EPEL

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
CMD ["build", "oraclelinux-8"]

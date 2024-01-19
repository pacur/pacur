FROM docker.io/oraclelinux:8
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y install dnf-plugins-core oracle-epel-release-el8

RUN dnf config-manager --set-enabled ol8_appstream
RUN dnf config-manager --set-enabled ol8_addons
RUN dnf config-manager --set-enabled ol8_codeready_builder
RUN dnf config-manager --set-enabled ol8_developer_EPEL

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
RUN echo "3f934f40ac360b9c01f616a9aa1796d227d8b0328bf64cb045c7b8c4ee9caea4 go1.21.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.21.6.linux-amd64.tar.gz
RUN rm -f go1.21.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-8"]

FROM docker.io/oraclelinux:9
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y install dnf-plugins-core oracle-epel-release-el9

RUN dnf config-manager --set-enabled ol9_appstream
RUN dnf config-manager --set-enabled ol9_addons
RUN dnf config-manager --set-enabled ol9_codeready_builder
RUN dnf config-manager --set-enabled ol9_developer_EPEL

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.24.3.linux-amd64.tar.gz
RUN echo "3333f6ea53afa971e9078895eaa4ac7204a8c6b5c68c10e6bc9a33e8e391bdd8 go1.24.3.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.24.3.linux-amd64.tar.gz
RUN rm -f go1.24.3.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-9"]

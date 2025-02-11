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

RUN wget https://go.dev/dl/go1.23.6.linux-amd64.tar.gz
RUN echo "9379441ea310de000f33a4dc767bd966e72ab2826270e038e78b2c53c2e7802d go1.23.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.23.6.linux-amd64.tar.gz
RUN rm -f go1.23.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-9"]

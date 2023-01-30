FROM docker.io/oraclelinux:9
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install dnf-plugins-core oracle-epel-release-el9

sudo dnf config-manager --set-enabled ol9_appstream
sudo dnf config-manager --set-enabled ol9_addons
sudo dnf config-manager --set-enabled ol9_codeready_builder
sudo dnf config-manager --set-enabled ol9_developer_EPEL

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.19.3.linux-amd64.tar.gz
RUN echo "74b9640724fd4e6bb0ed2a1bc44ae813a03f1e72a4c76253e2d5c015494430ba go1.19.3.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.19.3.linux-amd64.tar.gz
RUN rm -f go1.19.3.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-9"]

FROM docker.io/oraclelinux:10
LABEL maintainer="Pacur <contact@pacur.org>"

RUN dnf -y install oracle-epel-release-el10

RUN dnf config-manager --set-enabled ol10_appstream
RUN dnf config-manager --set-enabled ol10_addons
RUN dnf config-manager --set-enabled ol10_codeready_builder
RUN dnf config-manager --set-enabled ol10_u0_developer_EPEL

RUN dnf -y upgrade
RUN dnf -y groupinstall 'Development Tools'
RUN dnf -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.24.6.linux-amd64.tar.gz
RUN echo "bbca37cc395c974ffa4893ee35819ad23ebb27426df87af92e93a9ec66ef8712 go1.24.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.24.6.linux-amd64.tar.gz
RUN rm -f go1.24.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-10"]

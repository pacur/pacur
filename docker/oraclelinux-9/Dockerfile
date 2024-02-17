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

RUN wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
RUN echo "f6c8a87aa03b92c4b0bf3d558e28ea03006eb29db78917daec5cfb6ec1046265 go1.22.0.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.22.0.linux-amd64.tar.gz
RUN rm -f go1.22.0.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-9"]

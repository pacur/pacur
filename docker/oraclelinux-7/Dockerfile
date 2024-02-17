FROM docker.io/oraclelinux:7
LABEL maintainer="Pacur <contact@pacur.org>"

COPY RPM-GPG-KEY-endpoint-7 /etc/pki/rpm-gpg/RPM-GPG-KEY-endpoint-7
COPY endpoint.repo /etc/yum.repos.d/endpoint.repo
RUN yum -y install oracle-epel-release-el7

RUN yum -y install yum-utils
RUN yum-config-manager --enable ol7_developer
RUN yum-config-manager --enable ol7_developer_EPEL
RUN yum-config-manager --enable ol7_optional_latest
RUN yum-config-manager --enable ol7_addons

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
CMD ["build", "oraclelinux-7"]

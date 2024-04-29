FROM docker.io/amazonlinux:2023
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial
RUN yum -y install --allowerasing gnupg2

RUN wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
RUN echo "5901c52b7a78002aeff14a21f93e0f064f74ce1360fce51c6ee68cd471216a17 go1.22.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.22.2.linux-amd64.tar.gz
RUN rm -f go1.22.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-2023"]

FROM docker.io/fedora:36
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect gpg rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz
RUN echo "5e8c5a74fe6470dd7e055a461acda8bb4050ead8c2df70f227e3ff7d8eb7eeb6 go1.19.2.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.19.2.linux-amd64.tar.gz
RUN rm -f go1.19.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "fedora-36"]

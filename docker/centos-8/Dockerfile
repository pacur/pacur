FROM docker.io/centos:8
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

RUN yum -y install yum-utils

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.17.7.linux-amd64.tar.gz
RUN echo "02b111284bedbfa35a7e5b74a06082d18632eff824fd144312f6063943d49259 go1.17.7.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.7.linux-amd64.tar.gz
RUN rm -f go1.17.7.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos-8"]

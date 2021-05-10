FROM docker.io/centos:8
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

RUN yum -y install yum-utils

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz
RUN echo "7154e88f5a8047aad4b80ebace58a059e36e7e2e4eb3b383127a28c711b4ff59 go1.16.4.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.16.4.linux-amd64.tar.gz
RUN rm -f go1.16.4.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos-8"]

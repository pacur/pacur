FROM centos:8
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

RUN yum -y install yum-utils
RUN yum-config-manager --enable AppStream
RUN yum-config-manager --enable PowerTools

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://golang.org/dl/go1.15.5.linux-amd64.tar.gz
RUN echo "9a58494e8da722c3aef248c9227b0e9c528c7318309827780f16220998180a0d go1.15.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.15.5.linux-amd64.tar.gz
RUN rm -f go1.15.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos-8"]

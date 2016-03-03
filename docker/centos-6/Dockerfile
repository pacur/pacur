FROM centos:6
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install wget
RUN wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
RUN rpm -i epel-release-latest-6.noarch.rpm
RUN rm -f epel-release-latest-6.noarch.rpm
RUN yum -y upgrade
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.6.linux-amd64.tar.gz
RUN rm -f go1.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos-6"]

FROM centos:6
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install wget
RUN wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-6.noarch.rpm
RUN rpm -i epel-release-latest-6.noarch.rpm
RUN yum -y install tar make automake gcc wget zip git mercurial golang

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos"]

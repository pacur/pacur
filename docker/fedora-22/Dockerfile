FROM fedora:22
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y upgrade
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial golang

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "fedora"]

FROM amazonlinux:1
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install epel-release
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
RUN echo "2a3fdabf665496a0db5f41ec6af7a9b15a49fbe71a85a50ca38b1f13a103aeec go1.12.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.12.1.linux-amd64.tar.gz
RUN rm -f go1.12.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-1"]

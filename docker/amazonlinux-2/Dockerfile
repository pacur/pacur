FROM docker.io/amazonlinux:2
MAINTAINER Pacur <contact@pacur.org>

RUN rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
RUN echo "bd78114b0d441b029c8fe0341f4910370925a4d270a6a590668840675b0c653e go1.17.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.17.5.linux-amd64.tar.gz
RUN rm -f go1.17.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE off

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-2"]

FROM centos:7
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install epel-release
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
RUN echo "eb61005f0b932c93b424a3a4eaa67d72196c79129d9a3ea8578047683e2c80d5 go1.15.4.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.15.4.linux-amd64.tar.gz
RUN rm -f go1.15.4.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos-7"]

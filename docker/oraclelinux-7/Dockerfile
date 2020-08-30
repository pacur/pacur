FROM oraclelinux:7
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install yum-utils
RUN yum-config-manager --enable ol7_developer
RUN yum-config-manager --enable ol7_developer_epel
RUN yum-config-manager --enable ol7_optional_latest
RUN yum-config-manager --enable ol7_addons

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://dl.google.com/go/go1.15.linux-amd64.tar.gz
RUN echo "2d75848ac606061efe52a8068d0e647b35ce487a15bb52272c427df485193602 go1.15.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.15.linux-amd64.tar.gz
RUN rm -f go1.15.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-7"]

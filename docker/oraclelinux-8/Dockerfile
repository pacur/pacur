FROM oraclelinux:8
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

RUN yum -y install yum-utils
RUN yum-config-manager --enable ol8_appstream
RUN yum-config-manager --enable ol8_addons
RUN yum-config-manager --enable ol8_codeready_builder

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://dl.google.com/go/go1.14.5.linux-amd64.tar.gz
RUN echo "82a1b84f16858db03231eb201f90cce2a991078dda543879b87e738e2586854b go1.14.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.14.5.linux-amd64.tar.gz
RUN rm -f go1.14.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-8"]

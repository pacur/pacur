FROM amazonlinux:2
MAINTAINER Pacur <contact@pacur.org>

RUN rpm -Uvh https://mirror.sfo12.us.leaseweb.net/epel/7/x86_64/Packages/e/epel-release-7-11.noarch.rpm
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.9.2.linux-amd64.tar.gz
RUN rm -f go1.9.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-2"]

FROM fedora:24
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y upgrade
RUN yum -y install tar expect gpg rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.7.4.linux-amd64.tar.gz
RUN rm -f go1.7.4.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "fedora-24"]

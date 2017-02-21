FROM fedora:24
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y upgrade
RUN yum -y install tar expect gpg rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.8.src.tar.gz
RUN tar -C /usr/local -xf go1.8.src.tar.gz
RUN rm -f go1.8.src.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "fedora-24"]

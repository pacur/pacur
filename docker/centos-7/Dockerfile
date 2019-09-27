FROM centos:7
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y install epel-release
RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
RUN echo "94f874037b82ea5353f4061e543681a0e79657f787437974214629af8407d124 go1.13.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.13.1.linux-amd64.tar.gz
RUN rm -f go1.13.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "centos-7"]

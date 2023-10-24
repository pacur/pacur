FROM docker.io/fedora:38
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect gpg rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
RUN echo "1241381b2843fae5a9707eec1f8fb2ef94d827990582c7c7c32f5bdfbfd420c8 go1.21.3.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.21.3.linux-amd64.tar.gz
RUN rm -f go1.21.3.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "fedora-38"]

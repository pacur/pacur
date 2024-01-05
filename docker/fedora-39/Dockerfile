FROM docker.io/fedora:39
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect gpg rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
RUN echo "e2bc0b3e4b64111ec117295c088bde5f00eeed1567999ff77bc859d7df70078e go1.21.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.21.5.linux-amd64.tar.gz
RUN rm -f go1.21.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "fedora-39"]

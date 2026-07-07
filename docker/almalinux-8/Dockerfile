FROM docker.io/almalinux:8
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

RUN yum -y install dnf-plugins-core
RUN dnf config-manager --set-enabled powertools

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.26.5.linux-amd64.tar.gz
RUN echo "5c2c3b16caefa1d968a94c1daca04a7ca301a496d9b086e17ad77bb81393f053 go1.26.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.26.5.linux-amd64.tar.gz
RUN rm -f go1.26.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@bee0432f4c313ddf538afabbe0f9a3c87e0f0119

ENTRYPOINT ["pacur"]
CMD ["build", "almalinux-8"]

FROM docker.io/almalinux:8
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-8.noarch.rpm

RUN yum -y install dnf-plugins-core
RUN dnf config-manager --set-enabled powertools

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
RUN echo "904b924d435eaea086515bc63235b192ea441bd8c9b198c507e85009e6e4c7f0 go1.22.5.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.22.5.linux-amd64.tar.gz
RUN rm -f go1.22.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "almalinux-8"]

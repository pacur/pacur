FROM docker.io/almalinux:10
LABEL maintainer="Pacur <contact@pacur.org>"

RUN dnf -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-10.noarch.rpm

RUN dnf config-manager --set-enabled crb

RUN dnf -y upgrade
RUN dnf -y groupinstall 'Development Tools'
RUN dnf -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz
RUN echo "7716a0d940a0f6ae8e1f3b3f4f36299dc53e31b16840dbd171254312c41ca12e go1.25.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.25.1.linux-amd64.tar.gz
RUN rm -f go1.25.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "almalinux-10"]

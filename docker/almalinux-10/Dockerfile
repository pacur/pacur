FROM docker.io/almalinux:10
LABEL maintainer="Pacur <contact@pacur.org>"

RUN dnf -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-10.noarch.rpm

RUN dnf config-manager --set-enabled crb

RUN dnf -y upgrade
RUN dnf -y groupinstall 'Development Tools'
RUN dnf -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.26.4.linux-amd64.tar.gz
RUN echo "1153d3d50e0ac764b447adfe05c2bcf08e889d42a02e0fe0259bd47f6733ad7f go1.26.4.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.26.4.linux-amd64.tar.gz
RUN rm -f go1.26.4.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "almalinux-10"]

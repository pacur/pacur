FROM docker.io/oraclelinux:10
LABEL maintainer="Pacur <contact@pacur.org>"

RUN dnf -y install oracle-epel-release-el10

RUN dnf config-manager --set-enabled ol10_appstream
RUN dnf config-manager --set-enabled ol10_addons
RUN dnf config-manager --set-enabled ol10_codeready_builder

RUN dnf -y upgrade
RUN dnf -y groupinstall 'Development Tools'
RUN dnf -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial

RUN wget https://go.dev/dl/go1.27.1.linux-amd64.tar.gz
RUN echo "63d339f0da5ab53635a56f2490a7984dfe12dfcff22ad749f63edaf590168445 go1.27.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.27.1.linux-amd64.tar.gz
RUN rm -f go1.27.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@bee0432f4c313ddf538afabbe0f9a3c87e0f0119

ENTRYPOINT ["pacur"]
CMD ["build", "oraclelinux-10"]

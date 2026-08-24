FROM docker.io/amazonlinux:2023
LABEL maintainer="Pacur <contact@pacur.org>"

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install tar expect rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git mercurial
RUN yum -y install --allowerasing gnupg2

RUN wget https://go.dev/dl/go1.27.0.linux-amd64.tar.gz
RUN echo "675c26c449cbb18fc24b74650de1eabbae6e16f64326fd85a283fb3b58280685 go1.27.0.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.27.0.linux-amd64.tar.gz
RUN rm -f go1.27.0.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@bee0432f4c313ddf538afabbe0f9a3c87e0f0119

ENTRYPOINT ["pacur"]
CMD ["build", "amazonlinux-2023"]

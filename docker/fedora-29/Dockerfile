FROM fedora:29
MAINTAINER Pacur <contact@pacur.org>

RUN yum -y upgrade
RUN yum -y groupinstall 'Development Tools'
RUN yum -y install --allowerasing expect
RUN yum -y install tar expect gpg rpm-build rpm-sign openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
RUN echo "66d83bfb5a9ede000e33c6579a91a29e6b101829ad41fffb5c5bb6c900e109d9 go1.12.7.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.12.7.linux-amd64.tar.gz
RUN rm -f go1.12.7.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "fedora-29"]

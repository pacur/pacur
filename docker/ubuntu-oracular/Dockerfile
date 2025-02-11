FROM docker.io/ubuntu:oracular
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.23.6.linux-amd64.tar.gz
RUN echo "9379441ea310de000f33a4dc767bd966e72ab2826270e038e78b2c53c2e7802d go1.23.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.23.6.linux-amd64.tar.gz
RUN rm -f go1.23.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-oracular"]

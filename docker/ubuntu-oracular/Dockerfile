FROM docker.io/ubuntu:oracular
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.24.6.linux-amd64.tar.gz
RUN echo "bbca37cc395c974ffa4893ee35819ad23ebb27426df87af92e93a9ec66ef8712 go1.24.6.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.24.6.linux-amd64.tar.gz
RUN rm -f go1.24.6.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN GOPROXY=direct go install github.com/pacur/pacur@latest

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-oracular"]

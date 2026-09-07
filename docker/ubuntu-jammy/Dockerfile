FROM docker.io/ubuntu:jammy
LABEL maintainer="Pacur <contact@pacur.org>"

RUN apt-get --assume-yes update
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes upgrade
RUN DEBIAN_FRONTEND=noninteractive apt-get --assume-yes install build-essential reprepro rsync wget zip git mercurial

RUN wget https://go.dev/dl/go1.27.1.linux-amd64.tar.gz
RUN echo "63d339f0da5ab53635a56f2490a7984dfe12dfcff22ad749f63edaf590168445 go1.27.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.27.1.linux-amd64.tar.gz
RUN rm -f go1.27.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin
ENV GO111MODULE on

RUN go install github.com/pacur/pacur@bee0432f4c313ddf538afabbe0f9a3c87e0f0119

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-jammy"]

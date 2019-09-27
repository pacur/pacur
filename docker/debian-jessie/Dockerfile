FROM debian:jessie
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git bzr mercurial

RUN wget https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
RUN echo "94f874037b82ea5353f4061e543681a0e79657f787437974214629af8407d124 go1.13.1.linux-amd64.tar.gz" | sha256sum -c -
RUN tar -C /usr/local -xf go1.13.1.linux-amd64.tar.gz
RUN rm -f go1.13.1.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "debian-jessie"]

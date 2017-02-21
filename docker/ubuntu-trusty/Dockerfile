FROM ubuntu:trusty
MAINTAINER Pacur <contact@pacur.org>

RUN apt-get --assume-yes update
RUN apt-get --assume-yes upgrade
RUN apt-get --assume-yes install build-essential reprepro rsync wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.8.src.tar.gz
RUN tar -C /usr/local -xf go1.8.src.tar.gz
RUN rm -f go1.8.src.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "ubuntu-trusty"]

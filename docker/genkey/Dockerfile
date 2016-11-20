FROM pritunl/archlinux
MAINTAINER Pacur <contact@pacur.org>

RUN pacman -S --noconfirm go wget git

ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur", "genkey"]

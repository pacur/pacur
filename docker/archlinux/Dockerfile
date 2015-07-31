FROM pritunl/archlinux
MAINTAINER Pacur <contact@pacur.org>

RUN pacman -S --noconfirm base-devel go git mercurial wget rsync tar zip sudo

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "archlinux"]

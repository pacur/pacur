FROM base/archlinux
MAINTAINER Pacur <contact@pacur.org>

RUN pacman -S --noconfirm base-devel go git bzr mercurial wget rsync tar zip sudo
RUN ln -s -f /usr/bin/pinentry-curses /usr/bin/pinentry

ENV GOPATH /go
ENV PATH $PATH:/go/bin

RUN go get github.com/pacur/pacur

RUN sed -i 's|bsdtar -xf "$dbfile" -C "$tmpdir/$repo"|tar -xf "$dbfile" -C "$tmpdir/$repo"|g' /usr/bin/repo-add

ENTRYPOINT ["pacur"]
CMD ["build", "archlinux"]

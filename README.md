[![pacur](https://raw.githubusercontent.com/pacur/pacur-artwork/master/pacur.png)](http://pacur.org)

# pacur: simple packaging (beta)

[![Docker Repository](https://img.shields.io/badge/docker-archlinux-1794d1.svg?style=flat "Docker Repository")](https://registry.hub.docker
.com/u/pacur/archlinux/)
[![Docker Repository](https://img.shields.io/badge/docker-centos--6-10233f.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/centos-6/)
[![Docker Repository](https://img.shields.io/badge/docker-centos--7-10233f.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/centos-7/)
[![Docker Repository](https://img.shields.io/badge/docker-debian--jessie-d70a53.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/debian-jessie/)
[![Docker Repository](https://img.shields.io/badge/docker-debian--wheezy-d70a53.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/debian-wheezy/)


[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--precise-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-precise/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--trusty-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-trusty/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--vivid-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-vivid/)
[![Docker Repository](https://img.shields.io/badge/docker-ubuntu--wily-dd4814.svg?style=flat "Docker Repository")](https://registry.hub.docker.com/u/pacur/ubuntu-wily/)

Pacur allows building packages for several package formats and linux
distributions. Currently DEB, RPM and PKG packages are available for
ArchLinux, CentOS 6, CentOS 7, Debian 7, Debian 8, Ubuntu 12.04, Ubuntu 14.04,
Ubuntu 15.04 and Ubuntu 15.10. Builds are done on Docker containers without
needing to setup any virtual machines or install any software other then
Docker. All packages are built using a simple format that is similar to
[PKGBUILD](https://wiki.archlinux.org/index.php/PKGBUILD) from ArchLinux.
Each distribution is different and will still require a separate PKGBUILD for
each distribution but a consistent build process and format can be used for
all builds. Docker only supports 64 bit containers, pacur can't be used to
build packages 32 bit packages. Pacur will also create a DEB, APT and PKG
signed repository that can be used on ArchLinux, CentOS, Debian and Ubuntu to
distribute the packages.

### format

```
key="example string"
key=`example "quoted" string`
key=("list with one element")
key=(
    "list with"
    "multiple elements"
)
key="example ${variable} string"
```

### builtin variables

| key | value |
| --- | ----- |
| `${srcdir}` | `Source` directory where all sources are downloaded and extracted |
| `${pkgdir}` | `Package` directory for the root of the package |

### spec

| key | type | value |
| --- | ---- | ----- |
| `pkgname` | `string` | Package name |
| `pkgver` | `string` | Package version |
| `pkgrel` | `string` | Package release number |
| `pkgdesc` | `string` | Short package description |
| `pkgdesclong` | `list` | List of lines for package description |
| `maintainer` | `string` | Package maintainer |
| `arch` | `string` | Package architecture, can be `all` or `amd64` |
| `license` | `list` | List of licenses for packaged software |
| `section` | `string` | Section for package. Built in sections available:<br> `admin`<br> `localization`<br> `mail`<br> `comm`<br> `math`<br> `database`<br> `misc`<br> `debug`<br> `net`<br> `news`<br> `devel`<br> `doc`<br> `editors`<br> `electronics`<br> `embedded`<br> `fonts`<br> `games`<br> `science`<br> `shells`<br> `sound`<br> `graphics`<br> `text`<br> `httpd`<br> `vcs`<br> `interpreters`<br> `video`<br> `web`<br> `kernel`<br> `x11`<br> `libdevel`<br> `libs` |
| `priority` | `string` | Package priority, only used for debian packages |
| `url` | `string` | Package url |
| `depends` | `list` | List of package dependencies |
| `optdepends` | `list` | List of package optional dependencies |
| `makedepends` | `list` | List of package build dependencies |
| `provides` | `list` | List of packages provided |
| `conflicts` | `list` | List of packages conflicts |
| `sources` | `list` | List of packages sources. Sources can be url or paths that are relative to the PKGBUILD |
| `hashsums` | `list` | List of `md5`/`sha1`/`sha256`/`sha512` hex hashes for sources, hash type is determined by the length of the hash. Use `skip` to ignore hash check |
| `backup` | `list` | List of config files that shouldn't be overwritten on upgrades |
| `build` | `func` | Function to build the source, starts in srcdir |
| `package` | `func` | Function to package the source into the pkgdir, starts in srcdir |
| `preinst` | `func` | Function to run before installing |
| `postint` | `func` | Function to run after installing |
| `prerm` | `func` | Function to run before removing |
| `postrm` | `func` | Function to run after removing |

### example

First create a directory for the PKGBUILD file. This directory should only
contain the PKGBUILD file and any other files needed such as patches. Then
create a PKGBUILD the package directory. After creating the PKGBUILD build
the package with docker.

```
$ mkdir httpserver
$ cd httpserver
$ nano PKGBUILD
$ docker run --rm -t -v `pwd`:/pacur pacur/ubuntu-trusty
```

```
pkgname="httpserver"
pkgver="1.0"
pkgrel="1"
pkgdesc="Http file server written with Go"
pkgdesclong=(
    "Quick http file server written with Go"
    "using directory listing similar to apache"
)
maintainer="Pacur <contact@pacur.com>"
arch="all"
license=("GPLv3")
section="utils"
priority="optional"
url="https://github.com/pacur/${pkgname}"
sources=(
    "${url}/archive/${pkgver}.tar.gz"
)
hashsums=(
    "8ad577961ebe2e4377edbdda70459774"
)

build() {
    mkdir -p "go/src"
    export GOPATH="${srcdir}/go"
    mv "${pkgname}-${pkgver}" "go/src"
    cd "go/src/${pkgname}-${pkgver}"
    go get
    go build -a
}

package() {
    cd "${srcdir}/go/src/${pkgname}-${pkgver}"
    mkdir -p "${pkgdir}/usr/bin"
    cp ${pkgname}-${pkgver} ${pkgdir}/usr/bin/${pkgname}
}
```

### project example

A project can be created with the cli tools which can be installed using
go get. After creating a project a PKGBUILD file must be added to each release
directory inside the packages directory. Then the packages can be built and
added to the repo. An example project is available in the example directory.
The `pull` command should be run before all builds to update the docker images
used for builds.

```
$ go get github.com/pacur/pacur
$ cd example
$ pacur project init
$ pacur project pull
$ pacur project build
$ pacur project repo
$ go get github.com/pacur/httpserver
$ cd mirror
$ httpserver --port 80
```

After the repo has been created and is hosted on a server the following
commands can be used to add the repo to the package manager for yum and apt.
For the debian repo the `jessie` should be replaced with the debian/ubuntu
release name. The `pacur` repo name and filenames can be change to suite the
name of your software.

```
$ nano /etc/yum.repos.d/pacur.repo
[pacur]
name=Pacur Repository
baseurl=http://HTTP_SERVER_IP/yum/centos/7/
gpgcheck=0
enabled=1
$ yum install httpserver
```

```
$ nano /etc/apt/sources.list.d/pacur.list
deb http://HTTP_SERVER_IP/apt jessie main
$ apt-get update
$ apt-get install httpserver
```

### signing

Packages in the repository can also be signed by adding a `sign.key` in the
project directory or the `example` directory. The signing key cannot use a
passphrase. To export a key first get the key id then export the key with the
commands below.

```
$ gpg --list-secret-keys
$ gpg -a --export-secret-keys CF8E292A > sign.key
```

[![pacur](https://raw.githubusercontent.com/pacur/pacur-artwork/master/pacur.png)](http://pacur.org)

# pacur: simple packaging

![Arch Linux](https://img.shields.io/badge/podman-archlinux-1794d1.svg?style=flat "Arch Linux")


![AlmaLinux 8](https://img.shields.io/badge/podman-almalinux--8-0f4266.svg?style=flat "AlmaLinux 8")
![AlmaLinux 9](https://img.shields.io/badge/podman-almalinux--9-0f4266.svg?style=flat "AlmaLinux 9")


![Amazon Linux 2](https://img.shields.io/badge/podman-amazonlinux--2-faaf34.svg?style=flat "Amazon Linux 2")
![Amazon Linux 2023](https://img.shields.io/badge/podman-amazonlinux--2023-faaf34.svg?style=flat "Amazon Linux 2023")


![Debian Buster](https://img.shields.io/badge/podman-debian--buster-d70a53.svg?style=flat "Debian Buster")
![Debian Bullseye](https://img.shields.io/badge/podman-debian--bullseye-d70a53.svg?style=flat "Debian Bullseye")
![Debian Bookworm](https://img.shields.io/badge/podman-debian--bookworm-d70a53.svg?style=flat "Debian Bookworm")


![Oracle Linux 7](https://img.shields.io/badge/podman-oraclelinux--7-f82200.svg?style=flat "Oracle Linux 7")
![Oracle Linux 8](https://img.shields.io/badge/podman-oraclelinux--8-f82200.svg?style=flat "Oracle Linux 8")
![Oracle Linux 9](https://img.shields.io/badge/podman-oraclelinux--9-f82200.svg?style=flat "Oracle Linux 9")


![RockyLinux 8](https://img.shields.io/badge/podman-rockylinux--8-10b981.svg?style=flat "RockyLinux 8")
![RockyLinux 9](https://img.shields.io/badge/podman-rockylinux--9-10b981.svg?style=flat "RockyLinux 9")


![Fedora 40](https://img.shields.io/badge/podman-fedora--40-294172.svg?style=flat "Fedora 40")
![Fedora 41](https://img.shields.io/badge/podman-fedora--41-294172.svg?style=flat "Fedora 41")
![Fedora 42](https://img.shields.io/badge/podman-fedora--42-294172.svg?style=flat "Fedora 42")


![Ubuntu Bionic](https://img.shields.io/badge/podman-ubuntu--bionic-dd4814.svg?style=flat "Ubuntu Bionic")
![Ubuntu Focal](https://img.shields.io/badge/podman-ubuntu--focal-dd4814.svg?style=flat "Ubuntu Focal")
![Ubuntu Jammy](https://img.shields.io/badge/podman-ubuntu--jammy-dd4814.svg?style=flat "Ubuntu Jammy")
![Ubuntu Noble](https://img.shields.io/badge/podman-ubuntu--noble-dd4814.svg?style=flat "Ubuntu Noble")
![Ubuntu Oracular](https://img.shields.io/badge/podman-ubuntu--oracular-dd4814.svg?style=flat "Ubuntu Oracular")
![Ubuntu Plucky](https://img.shields.io/badge/podman-ubuntu--plucky-dd4814.svg?style=flat "Ubuntu Plucky")

Pacur allows building packages for multiple linux distributions with a
consistent package spec format. Currently `deb`, `rpm` and `pacman` packages
are available for several linux distributions. Builds are done on podman
containers without needing to setup any virtual machines or install any
software other than podman. All packages are built using a simple format that
is similar to [PKGBUILD](https://wiki.archlinux.org/index.php/PKGBUILD) from
ArchLinux. Each distribution is different and will still require different
build instructions, but a consistent build process and format can be used for
all builds. podman only supports 64-bit containers, pacur can't be used to
build 32-bit packages. Pacur will also create a `deb`, `rpm` and
`pacman` signed repository that can be used on ArchLinux, CentOS, Fedora,
Debian and Ubuntu to distribute the packages. A
[tutorial](https://medium.com/@zachhuff386/pacur-tutorial-9848b774c84a)
on creating a project is available on medium.

### initialize

It is recommended to build the podman images locally instead of pulling each
image from the podman Hub. Currently the podman Hub images are not
maintained. A script is located in the podman directory to assist with this.
Always run the `clean.sh` script to clear any existing pacur images. Building
the images can take several hours.

```
go install github.com/pacur/pacur@latest
cd "$(ls -d ~/go/pkg/mod/github.com/pacur/pacur@*/podman/ | sort -V | tail -n 1)"
sh clean.sh
sh build.sh
```

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
key:ubuntu="this will apply only to ubuntu builds"
```

### builtin variables

| key         | value |
|-------------| ----- |
| `${srcdir}` | `Source` directory where all sources are downloaded and extracted |
| `${pkgdir}` | `Package` directory for the root of the package |

### spec

| key          | type | value |
|--------------| ---- | ----- |
| `targets`    | `list` | List of build targets only used for projects. Prefix a `!` to ignore target. |
| `pkgname`    | `string` | Package name |
| `pkgver`     | `string` | Package version |
| `pkgrel`     | `string` | Package release number |
| `pkgdesc`    | `string` | Short package description |
| `pkgdesclong` | `list` | List of lines for package description |
| `maintainer` | `string` | Package maintainer |
| `arch`       | `string` | Package architecture, can be `all` or `amd64` |
| `license`    | `list` | List of licenses for packaged software |
| `section`    | `string` | Section for package. Built in sections available:<br> `admin`<br> `localization`<br> `mail`<br> `comm`<br> `math`<br> `database`<br> `misc`<br> `debug`<br> `net`<br> `news`<br> `devel`<br> `doc`<br> `editors`<br> `electronics`<br> `embedded`<br> `fonts`<br> `games`<br> `science`<br> `shells`<br> `sound`<br> `graphics`<br> `text`<br> `httpd`<br> `vcs`<br> `interpreters`<br> `video`<br> `web`<br> `kernel`<br> `x11`<br> `libdevel`<br> `libs` |
| `priority`   | `string` | Package priority, only used for debian packages |
| `url`        | `string` | Package url |
| `rpmopts`    | `list` | List of lines to add to RPM spec |
| `depends`    | `list` | List of package dependencies |
| `optdepends` | `list` | List of package optional dependencies |
| `makedepends` | `list` | List of package build dependencies |
| `provides`   | `list` | List of packages provided |
| `conflicts`  | `list` | List of packages conflicts |
| `sources`    | `list` | List of packages sources. Sources can be url or paths that are relative to the PKGBUILD |
| `hashsums`   | `list` | List of `md5`/`sha1`/`sha256`/`sha512` hex hashes for sources, hash type is determined by the length of the hash. Use `skip` to ignore hash check |
| `backup`     | `list` | List of config files that shouldn't be overwritten on upgrades |
| `build`      | `func` | Function to build the source, starts in srcdir |
| `package`    | `func` | Function to package the source into the pkgdir, starts in srcdir |
| `preinst`    | `func` | Function to run before installing |
| `postinst`   | `func` | Function to run after installing |
| `prerm`      | `func` | Function to run before removing |
| `postrm`     | `func` | Function to run after removing |

### build targets

| target             | value                    |
|--------------------|--------------------------|
| `archlinux`        | All archlinux releases   |
| `almalinux`        | All almalinux releases   |
| `amazonlinux`      | All amazonlinux releases |
| `debian`           | All debian releases      |
| `fedora`           | All fedora releases      |
| `oraclelinux`      | All oraclelinux releases |
| `rockylinux`       | All rockylinux releases  |
| `ubuntu`           | All ubuntu releases      |
| `almalinux-8`      | AlmaLinux 8              |
| `almalinux-9`      | AlmaLinux 9              |
| `amazonlinux-2`    | Amazonlinux 2            |
| `amazonlinux-2023` | Amazonlinux 2023         |
| `debian-buster`    | Debian buster            |
| `debian-bullseye`  | Debian bullseye          |
| `debian-bookworm`  | Debian bookworm          |
| `fedora-40`        | Fedora 40                |
| `fedora-41`        | Fedora 41                |
| `fedora-42`        | Fedora 42                |
| `oraclelinux-7`    | Oraclelinux 7            |
| `oraclelinux-8`    | Oraclelinux 8            |
| `oraclelinux-9`    | Oraclelinux 9            |
| `rockylinux-8`     | RockyLinux 8             |
| `rockylinux-9`     | RockyLinux 9             |
| `ubuntu-bionic`    | Ubuntu bionic            |
| `ubuntu-focal`     | Ubuntu focal             |
| `ubuntu-jammy`     | Ubuntu jammy             |
| `ubuntu-noble`     | Ubuntu noble             |
| `ubuntu-oracular`  | Ubuntu oracular          |
| `ubuntu-plucky`    | Ubuntu plucky            |

### directives

| directive          | value                    |
|--------------------|--------------------------|
| `apt`              | All deb packages         |
| `pacman`           | All pkg packages         |
| `yum`              | All rpm packages         |
| `archlinux`        | All archlinux releases   |
| `almalinux`        | All almalinux releases   |
| `amazonlinux`      | All amazonlinux releases |
| `debian`           | All debian releases      |
| `fedora`           | All fedora releases      |
| `oraclelinux`      | All oraclelinux releases |
| `rockylinux`       | All rockylinux releases  |
| `ubuntu`           | All ubuntu releases      |
| `almalinux-8`      | AlmaLinux 8              |
| `almalinux-9`      | AlmaLinux 9              |
| `amazonlinux-2`    | Amazonlinux 2            |
| `amazonlinux-2023` | Amazonlinux 2023         |
| `debian-buster`    | Debian buster            |
| `debian-bullseye`  | Debian bullseye          |
| `debian-bookworm`  | Debian bookworm          |
| `fedora-40`        | Fedora 40                |
| `fedora-41`        | Fedora 41                |
| `fedora-42`        | Fedora 42                |
| `oraclelinux-7`    | Oraclelinux 7            |
| `oraclelinux-8`    | Oraclelinux 8            |
| `oraclelinux-9`    | Oraclelinux 9            |
| `rockylinux-8`     | RockyLinux 8             |
| `rockylinux-9`     | RockyLinux 9             |
| `ubuntu-bionic`    | Ubuntu bionic            |
| `ubuntu-focal`     | Ubuntu focal             |
| `ubuntu-jammy`     | Ubuntu jammy             |
| `ubuntu-noble`     | Ubuntu noble             |
| `ubuntu-oracular`  | Ubuntu oracular          |
| `ubuntu-plucky`    | Ubuntu plucky            |

Directives are used to specify variables that only apply to a limited set of
build targets. All variables can use directives including user defined
variables. To use directives include the directive after a
variable separated by a colon such as
`pkgdesc:ubuntu="This description will only apply to Ubuntu packages"`.
The directives above are sorted from lowest to the highest priority.

### example

First create a directory for the PKGBUILD file. This directory should only
contain the PKGBUILD file and any other files needed such as patches. Then
create a PKGBUILD the package directory. After creating the PKGBUILD build
the package with podman.

```
$ mkdir httpserver
$ cd httpserver
$ nano PKGBUILD
$ podman run --rm -t -v `pwd`:/pacur pacur/ubuntu-xenial
```

```
targets=(
    "archlinux"
    "oraclelinux"
    "debian"
    "ubuntu"
)
pkgname="httpserver"
pkgver="1.0"
pkgrel="1"
pkgdesc="Http file server written with Go"
pkgdesc:oraclelinux="Http file server written with Go for Oracle Linux"
pkgdesc:debian="Http file server written with Go for Debian"
pkgdesc:fedora="Http file server written with Go for Fedora"
pkgdesc:ubuntu="Http file server written with Go for Ubuntu"
pkgdesclong=(
    "Quick http file server written with Go"
    "using directory listing similar to apache"
)
rpmopts=(
    "AutoReqProv: no"
)
maintainer="Example <example@pacur.org>"
arch="all"
license=("GPLv3")
section="utils"
priority="optional"
url="https://github.com/pacur/${pkgname}"
sources=(
    "${url}/archive/${pkgver}.tar.gz"
)
hashsums=(
    "3548e1263a931b27970e190f04b74623"
)

build() {
    mkdir -p "go/src"
    export GOPATH="${srcdir}/go"
    mv "${pkgname}-${pkgver}" "go/src"
    cd "go/src/${pkgname}-${pkgver}"
    go install
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
`go install`. The packages can be built and added to the repo. An example
project is available in the example directory. The container build script
must first be run to build the container images for all distributions.
Directories in the podman directory can be removed to exclude distributions
that are not needed.

```
$ go install github.com/pacur/pacur@latest
$ cd "$(ls -d ~/go/pkg/mod/github.com/pacur/pacur@*/podman/ | sort -V | tail -n 1)"
$ sh build.sh
$ cd example
$ pacur project init
$ pacur project build
$ pacur project repo
$ go install github.com/pacur/httpserver@latest
$ cd mirror
$ httpserver --port 80
```

After the repo has been created and is hosted on a server the following
commands can be used to add the repo to the package manager for yum and apt.
For the debian repo the `bullseye` should be replaced with the debian/ubuntu
release name. The `pacur` repo name and filenames can be change to suite the
name of your software.

```
$ nano /etc/pacman.conf
[pacur]
Server = http://HTTP_SERVER_IP/arch

$ pacman-key --keyserver hkp://pgp.mit.edu -r KEYID
$ pacman-key --lsign-key KEYID
$ pacman -Sy
$ pacman -S httpserver
```

```
$ nano /etc/yum.repos.d/pacur.repo
[pacur]
name=Pacur Repository
baseurl=http://HTTP_SERVER_IP/yum/oraclelinux/7/
gpgcheck=1
enabled=1

$ gpg --keyserver hkp://pgp.mit.edu --recv-keys KEYID
$ gpg --armor --export KEYID > key.tmp; rpm --import key.tmp; rm -f key.tmp
$ yum install httpserver
```

```
$ nano /etc/apt/sources.list.d/pacur.list
deb http://HTTP_SERVER_IP/apt bullseye main

$ apt-key adv --keyserver hkp://pgp.mit.edu --recv KEYID
$ apt-get update
$ apt-get install httpserver
```

### signing

Packages in the repository can also be signed by adding a `sign.key` in the
project directory. The signing key cannot use a passphrase. To export a key
first get the key id then export the key with the commands below.

```
$ gpg --list-secret-keys
$ gpg -a --export-secret-keys KEYID > sign.key
```

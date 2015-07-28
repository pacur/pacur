[![pacur](https://raw.githubusercontent.com/pacur/pacur-artwork/master/pacur.png)](http://pacur.org)

# pacur: simple packaging (beta)

Pacur allows building packages for several package formats and linux
distributions. Currently DEB and RPM packages are available for CentOS 6,
CentOS 7, Debian 7, Debian 8, Ubuntu 12.04, Ubuntu 14.04, Ubuntu 15.04 and
Ubuntu 15.10. Builds are done on Docker containers without needing to setup
any virtual machines or install any software other then Docker. All packages
are built using a simple format that is similar to
[PKGBUILD](https://wiki.archlinux.org/index.php/PKGBUILD) from Arch Linux.
Each distribution is different and will still require a seperate PKGBUILD for
each distribution but a consistent build process and format can be used for
all builds. Docker only supports 64 bit containers, pacur can't be used to
build packages 32 bit packages.

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
| `${srcdir}` | `Source` directory where all sources are downloaded and
extracted |
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
| `arch` | `string` | Package architecture, only used for debian packages |
| `license` | `list` | List of licenses for packaged software |
| `section` | `string` | Section for package. Built in sections available:<br> `admin`<br> `localization`<br> `mail`<br> `comm`<br> `math`<br> `database`<br> `misc`<br> `debug`<br> `net`<br> `news`<br> `devel`<br> `doc`<br> `editors`<br> `electronics`<br> `embedded`<br> `fonts`<br> `games`<br> `science`<br> `shells`<br> `sound`<br> `graphics`<br> `text`<br> `httpd`<br> `vcs`<br> `interpreters`<br> `video`<br> `web`<br> `kernel`<br> `x11`<br> `libdevel`<br>
`libs` |
| `priority` | `string` | Package priority, only used for debian packages |
| `url` | `string` | Package url |
| `depends` | `list` | List of package dependencies |
| `optdepends` | `list` | List of package optional dependencies |
| `makedepends` | `list` | List of package build dependencies |
| `provides` | `list` | List of packages provided |
| `conflicts` | `list` | List of packages conflicts |
| `sources` | `list` | List of packages sources |
| `hashsums` | `list` | List of `md5`/`sha1`/`sha256`/`sha512` hex hashes for sources, hash type is determined by the length of the hash. Use `skip` to ignore hash check |
| `backup` | `list` | List of config files that shouldn't be overwritten on upgrades |
| `build` | `func` | Function to build the source, starts in srcdir |
| `package` | `func` | Function to package the source into the pkgdir, starts in srcdir |
| `preinst` | `func` | Function to run before installing |
| `postint` | `func` | Function to run after installing |
| `prerm` | `func` | Function to run before removing |
| `postrm` | `func` | Function to run after removing |

Copyright (c) 2015 Zachary Huff

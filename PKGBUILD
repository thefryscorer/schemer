# Maintainer: Jacqueline Leykam <jleykam@gmail.com>

pkgname=schemer
pkgver=1.0
pkgrel=1
pkgdesc="Utility to generate terminal colorschemes from images."
arch=('x86_64' 'i686' 'arm')
url="https://github.com/thefryscorer/schemer"
makedepends=('git' 'go' 'sdl' 'sdl_image')
options=('!strip' '!emptydirs')
_gourl=github.com/thefryscorer/schemer

build() {
  GOPATH="$srcdir" go get -fix -v -x ${_gourl}
}

package() {
  mkdir -p "$pkgdir/usr/bin"
  install -p -m755 "$srcdir/bin/"* "$pkgdir/usr/bin"
}

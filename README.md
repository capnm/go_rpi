[wiki](https://github.com/capnm/go_rpi/wiki/)
[Go](http://golang.org/)
[RPi](http://www.raspberrypi.org/)
# Let's Go and have some fun with the Raspberry Pi.

### Install
```
sudo apt-get install libjpeg8-dev
git clone git://github.com/capnm/go_rpi.git

cd go_rpi
export GOPATH=$(pwd)
```

# Examples
### OpenVG: For a test if everything works, just draw a blue circle on the screen.
Run `bin/circle` in the linux console (ctrl+alt+F1) or the (ssh) terminal.
You can rebuild the binary with `go install -v circle`.

### OpenVG: A simple clock displaying some hardware data.
Run `bin/clock`.
Rebuild it with `go install -v clock`

![clock](img/clock.png)

# Install a bleeding edge version of the Go language.
```
git clone -b gitfix git://github.com/capnm/golang.git
cd golang
```

#### Gotchas
* OpenVG doesn't work with Go 1.0.3 release (no cgo support).
* OpenVG doesn't work with Go tip (bug https://code.google.com/p/go/issues/detail?id=5227).

```
# A workaround: checkout the last good commit.
git checkout ae11ae9fd2

# Nuke all garbage.
git gc --prune
git clean -fd

echo "devel gae11ae9fd2" > VERSION

cd src
./make.bash or ./all.bash
```
and add `golang/bin` to your PATH

# Kudos 
* The Go OpenVG parts have been extracted from [ajstarks](https://github.com/ajstarks/openvg) repo.

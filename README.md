[wiki](https://github.com/capnm/go_rpi/wiki/)
[Go](http://golang.org/)
[RPi](http://www.raspberrypi.org/)
# Let's Go and have some fun with the Raspberry Pi.

### Install
```
sudo apt-get install libjpeg8-dev git
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


# Gotchas
* OpenVG doesn't work with Go 1.0.3 release (no cgo support).

# Install a bleeding edge version of the Go language.
```
git clone -b gitfix git://github.com/capnm/golang.git
cd golang

# Nuke all garbage.
git gc --prune
git clean -fd

cd src
./make.bash or ./all.bash
```

# Credits 
* The Go OpenVG parts have been extracted from [ajstarks](https://github.com/ajstarks/openvg) repo.

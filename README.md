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
error: failed to add service - already in use?
	edit /boot/config.txt; make sure that the minimum video ram size (gpu_mem) is 64mb

# Install the Go language, release 1.3 (~280MB).
```
git clone -b release-branch.go1.3 git://github.com/capnm/golang.git --depth 1
sudo mv golang /opt/
cd opt/golang/src
./make.bash 

echo 'export PATH=/opt/golang/bin:$PATH' >> /etc/bash.bashrc

reboot

```

# Credits 
* The Go OpenVG parts have been extracted from the [ajstarks openvg](https://github.com/ajstarks/openvg) repo.

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
### To check if everything works, run the blue circle program:
Run `bin/circle` in the linux console (ctrl+alt+F1) or the (ssh) terminal.
You can rebuild the binary with `go install -v circle`.

### A simple clock displaying hardware data:
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
cd /opt/golang/src
./make.bash 

sudo -i
echo 'export PATH=/opt/golang/bin:$PATH' >> /etc/bash.bashrc

reboot

```

# Credits 
* The original OpenVG code was copied from the [ajstarks openvg](https://github.com/ajstarks/openvg) repository.

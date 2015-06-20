[wiki](https://github.com/capnm/go_rpi/wiki/)
[Go](http://golang.org/)
[RPi](http://www.raspberrypi.org/)
# A simple 2D graphics system for the Raspberry Pi.

### At first install the Go language and the jpeg library.
(type that line by line in the terminal)
```
sudo apt-get install libjpeg8-dev git

wget https://github.com/golang/go/archive/go1.4.2.tar.gz
sudo tar xf go1.4.2.tar.gz -C /opt/
sudo mv /opt/go-go1.4.2 /opt/go
sudo chown -R $USER /opt/go

# compile Go
cd /opt/go/src
./make.bash 
# wait a few minutes

# set correct permissions and make the Go language available in the terminal 
sudo chown -R root:root /opt/go
sudo -i
cd /usr/bin
ln -s /opt/go/bin/go
ln -s /opt/go/bin/gofmt
exit

# and test Go:
go version
--> go version go1.4.2 linux/arm

```

### Install and use this repository
```
git clone git://github.com/capnm/go_rpi.git

cd go_rpi
export GOPATH=$(pwd)
```

# Examples
### To check if everything works, run the blue circle example:
In the terminal -- or over ssh or in the linux console (ctrl+alt+F1) -- go with `cd` to the `go_rpi` directory, setup the Go environment by typing `export GOPATH=$(pwd)` and run the example by typing `bin/circle`.

You can change the source code in the `src/circle` directory and after that you must rebuild the program with `go install -v circle`.

### An example of a continuously updated clock displaying some of the SoC hardware data:
Run `bin/clock`.
Rebuild it with `go install -v clock`

![clock](img/clock.png)


# Gotchas
You get `error: failed to add service - already in use?` or an another weird `EGL error` message:
* Edit the `/boot/config.txt` file. Make sure that the video RAM size isn't set below 64MB (`gpu_mem=64`).



# Credits 
* The OpenVG library was vendored from the [ajstarks openvg](https://github.com/ajstarks/openvg) repository.

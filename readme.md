# This is a program for creating a simulation for libre 2 for xdrip.
## Examples

### Build and run the examples on a raspberry pi


sudo apt-get install golang -y
mkdir go
mkdir go/src
cd go/src/

git clone https://github.com/tzachi-dar/gatt.git

mkdir -p github.com/paypal
cd github.com/paypal 
git clone https://github.com/paypal/gatt
cd ../../gatt


# Build the sample server.
go build examples/server.go
# Start the sample server.
sudo ./server

make sure you run this program from a place that the shell will not be 
killed with time (for example using VNC).

# How to tell xdrip to connect to the new server:
1) Make sure that xdrip is using fake data. you can do this by changing the function
  NFCReaderX::use_fake_de_data to always return true.
2) Compile xdrip and install it to a phone.
3) Scan any old libre sensor, after that it will connect to the libre 2 simulator.

 


Gatt is released under a [BSD-style license](./LICENSE.md).

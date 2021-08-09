# Wake-on-LAN
Wake-on-LAN (WOL)
- With password support
- Written in golang 

## Usage
```
Version:
  1.00

Usage:
  wakeonlan {Command} [Option]

Command:
  -m  <MAC Address> : set target machine's mac address
  -h                : show help
  -v                : show version

Option:
  -pw <Password>    : set magic packet password
  -a  <IP>          : set broadcast IP
  -p  <Port>        : set udp port

Example:
  1) wakeonlan -m 11-22-33-44-55-66
  2) wakeonlan -m 11-22-33-44-55-66 -pw AA-BB-CC-DD-EE-FF -a 192.168.1.255 -p 9
  3) wakeonlan -h
  4) wakeonlan -v
```

## Build
### Example
```sh
# dependence
git clone https://github.com/gek64/gek.git

git clone https://github.com/gek64/WakeOnLAN.git

cd WakeOnLAN

go build -v -trimpath -ldflags "-s -w"
```

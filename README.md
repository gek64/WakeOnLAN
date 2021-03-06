# Wake-on-LAN
- wake on lan
- with password support
- customizable broadcast ip address and port

## Usage
```
Usage:
  wakeonlan {Command} [Option] MAC_ADDRESS

Command:
  -h                : show help
  -v                : show version                                           
                                                                             
Option:                                                                      
  -pw <Password>    : set magic packet password                              
  -a  <IP>          : set broadcast IP                                       
  -p  <Port>        : set udp port                                           
                                                                             
Example:                                                                     
  1) wakeonlan 1A-2B-3C-4D-5E-6F                                             
  2) wakeonlan 1A-2B-3C-4D-5E-6F -pw AA-BB-CC-DD-EE-FF -a 192.168.1.255 -p 9 
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

## License

**GPL-3.0 License**

See `LICENSE` for details


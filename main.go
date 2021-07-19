package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	cliMAC      string
	cliPassword string
	cliAddr     string
	cliPort     int
	cliHelp     bool
	cliVersion  bool
)

func init() {
	flag.StringVar(&cliMAC, "m", "00-00-00-00-00-00", "set target machine's mac address")
	flag.StringVar(&cliPassword, "pw", "00-00-00-00-00-00", "set magic packet password")
	flag.StringVar(&cliAddr, "a", "255.255.255.255", "set broadcast IP")
	flag.IntVar(&cliPort, "p", 9, "set udp port")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `
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
  4) wakeonlan -v`
		fmt.Println(helpInfo)
	}

	// 如果无args,打印用法后退出
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(0)
	}

	// 打印帮助信息
	if cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		showVersion()
		os.Exit(0)
	}
}

func showVersion() {
	var versionInfo = `Changelog:
  1.00:
    - First release`
	fmt.Println(versionInfo)
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	if isFlagPassed("m") {
		magicPacket, err := ParseMagicPacket(cliMAC, cliPassword)
		if err != nil {
			log.Fatal(err)
		}

		err = magicPacket.Send(cliAddr, cliPort)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Magic packet send successfully!")
	}
}

package main

import (
	"flag"
	"fmt"
	"gek_flag"
	"log"
	"net"
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

	// 判断IP是否合法
	if ip := net.ParseIP(cliAddr); ip == nil {
		log.Fatalf("%s is not a valid host", cliAddr)
	}
	// 判断端口是否合法
	if cliPort < 0 || cliPort > 65535 {
		log.Fatalf("%d is not a valid port", cliPort)
	}

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `
Version:
  1.01

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

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
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
    - First release
  1.01:
    - Modify fmt.Errorf() to avoid failure if the password is not a string`
	fmt.Println(versionInfo)
}

func main() {
	// 处理 m 参数
	if gek_flag.IsFlagPassed("m") {
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

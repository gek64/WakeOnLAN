package main

import (
	"fmt"
	"net"
	"strconv"
)

/*
魔法包结构:
     6字节头            96字节(目标mac * 16)        4-6字节(密码)
FF FF FF FF FF FF    11 22 33 44 55 66 *16    00 00 00 00 00 00
*/

// MagicPacket 魔法包
type MagicPacket struct {
	header   []byte
	payload  []byte
	password []byte
}

// ParseMagicPacket MAC地址字符串转换为魔法包
func ParseMagicPacket(hardwareAddr string, password ...interface{}) (MagicPacket, error) {
	var mp MagicPacket

	mp.header = make([]byte, 6)
	mp.payload = make([]byte, 96)
	mp.password = make([]byte, 6)

	// header
	for i := 0; i < 6; i++ {
		mp.header[i] = 0xff
	}

	// payload
	mac, err := net.ParseMAC(hardwareAddr)
	if err != nil {
		return MagicPacket{}, err
	}
	for i := 0; i < 96; i = i + 6 {
		copy(mp.payload[i:], mac)
	}

	// password
	if len(password) != 0 {
		passwd, err := net.ParseMAC(password[0].(string))
		if err != nil {
			return MagicPacket{}, fmt.Errorf("invalid password %s", password[0])
		}

		for i := 0; i < 6; i++ {
			mp.password[i] = passwd[i]
		}
	}

	return mp, nil
}

// Unmarshal []byte转换为魔法包
func Unmarshal(mpb []byte) (MagicPacket, error) {
	var mp MagicPacket

	mp.header = make([]byte, 6)
	mp.payload = make([]byte, 96)
	mp.password = make([]byte, 6)

	if len(mpb) == 108 {
		copy(mp.header[:], mpb[:6])
		copy(mp.payload[:], mpb[6:102])
		copy(mp.password[:], mpb[102:])
	} else if len(mpb) == 102 {
		copy(mp.header[:], mpb[:6])
		copy(mp.payload[:], mpb[6:])
	} else {
		return mp, fmt.Errorf("invalid magic packet")
	}

	return mp, nil
}

// Marshal 魔法包转换为[]byte
func (mp MagicPacket) Marshal() []byte {
	// 拼接 header payload password
	return append(append(mp.header, mp.payload...), mp.password...)
}

// Send 发送魔法包
func (mp MagicPacket) Send(addr string, port int) error {
	// 建立udp连接
	conn, err := net.Dial("udp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	// 函数运行完成关闭连接
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	// 发送数据
	_, err = conn.Write(mp.Marshal())
	if err != nil {
		return err
	}

	return nil
}

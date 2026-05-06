// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

package mdns

import "net"

func mdnsInterfaces() ([]net.Interface, error) {
	return nil, errTinyGoUnsupported
}

func mdnsInterfaceAddrs(net.Interface) ([]net.Addr, error) {
	return nil, errTinyGoUnsupported
}

func mdnsListenUDP(string, *net.UDPAddr) (*net.UDPConn, error) {
	return nil, errTinyGoUnsupported
}

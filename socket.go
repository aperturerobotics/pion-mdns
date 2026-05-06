// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo

package mdns

import "net"

func mdnsInterfaces() ([]net.Interface, error) {
	return net.Interfaces()
}

func mdnsInterfaceAddrs(ifc net.Interface) ([]net.Addr, error) {
	return ifc.Addrs()
}

func mdnsListenUDP(network string, addr *net.UDPAddr) (*net.UDPConn, error) {
	return net.ListenUDP(network, addr)
}

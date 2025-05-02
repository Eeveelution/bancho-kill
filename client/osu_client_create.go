package client

import (
	"bancho-kill/packets"
	"net"
)

func CreateTcpOsuClient(addr string, version int, failFunc func(string), warnFunc func(string)) (client *OsuClient, err error) {
	resolvedAddr, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, resolvedAddr)

	if err != nil {
		return nil, err
	}

	return &OsuClient{
		conn:                conn,
		serverAddr:          addr,
		kind:                ClientKindTcp,
		clientVersion:       version,
		PacketIncomingQueue: make(chan packets.BanchoPacket, 128),
		PacketOutgoingQueue: make(chan []byte, 128),
	}, nil
}

func CreateHttpOsuClient(addr string, version int, failFunc func(string), warnFunc func(string)) (client *OsuClient, err error) {
	return &OsuClient{
		conn:                nil,
		serverAddr:          addr,
		kind:                ClientKindHttp,
		clientVersion:       version,
		PacketIncomingQueue: make(chan packets.BanchoPacket, 128),
		PacketOutgoingQueue: make(chan []byte, 128),
		PacketPollRate:      500,
	}, nil
}

func CreateOsuClient(kind ClientKind, version int, addr string, failFunc func(string), warnFunc func(string)) (*OsuClient, error) {
	if kind == ClientKindTcp {
		return CreateTcpOsuClient(addr, version, failFunc, warnFunc)
	} else {
		return CreateHttpOsuClient(addr, version, failFunc, warnFunc)
	}
}

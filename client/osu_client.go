package client

import (
	"bancho-kill/middleware"
	"net"
)

type ClientKind int8

const (
	ClientKindTcp  ClientKind = 0
	ClientKindHttp ClientKind = 1
)

type OsuClient struct {
	conn        *net.TCPConn
	banchoToken string
	serverAddr  string
	kind        ClientKind

	PacketMiddleware []middleware.Middleware
	DataMiddleware   []middleware.Middleware

	ClientVersion       int
	PacketIncomingQueue chan []byte
	PacketOutgoingQueue chan []byte
}

func CreateTcpOsuClient(addr string, version int) (client *OsuClient, err error) {
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
		ClientVersion:       version,
		PacketIncomingQueue: make(chan []byte),
		PacketOutgoingQueue: make(chan []byte),
	}, nil
}

func CreateHttpOsuClient(addr string, version int) (client *OsuClient, err error) {
	return &OsuClient{
		conn:                nil,
		serverAddr:          addr,
		kind:                ClientKindHttp,
		ClientVersion:       version,
		PacketIncomingQueue: make(chan []byte),
		PacketOutgoingQueue: make(chan []byte),
	}, nil
}

func CreateOsuClient(kind ClientKind, version int, addr string) (*OsuClient, error) {
	if kind == ClientKindTcp {
		return CreateTcpOsuClient(addr, version)
	} else {
		return CreateHttpOsuClient(addr, version)
	}
}

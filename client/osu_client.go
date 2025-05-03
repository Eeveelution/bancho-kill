package client

import (
	"bancho-kill/middleware"
	"bancho-kill/packets"
	"net"
	"time"
)

type ClientKind int8

const (
	ClientKindTcp  ClientKind = 0
	ClientKindHttp ClientKind = 1
)

type OsuClient struct {
	conn            *net.TCPConn
	banchoToken     string
	serverAddr      string
	tcpAddr         *net.TCPAddr
	kind            ClientKind
	clientVersion   int
	continueRunning bool

	PacketReaderMiddleware middleware.PacketMiddleware
	PacketMiddleware       []middleware.PacketMiddleware
	DataMiddleware         []middleware.DataMiddleware

	PacketHistory       []packets.BanchoPacket
	PacketIncomingQueue chan packets.BanchoPacket

	PacketOutgoingQueue chan []byte
	PacketPollRate      time.Duration

	KillSignal chan struct{}

	TestFailFunc func(string)
	TestWarnFunc func(string)

	OwnUserData           UserData
	PresentUsers          []UserData
	JoinedChannels        []ChatChannel
	BanchoProtocolVersion int32
}

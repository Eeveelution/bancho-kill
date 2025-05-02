package middleware

import "bancho-kill/packets"

type PacketMiddleware interface {
	Process(packet packets.BanchoPacket) packets.BanchoPacket
}

type DataMiddleware interface {
	Process(data []byte) []byte
}

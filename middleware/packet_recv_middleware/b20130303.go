package packet_recv_middleware

import "bancho-kill/packets"

type ReceiverMiddleware_b20130303 struct{}

func (ReceiverMiddleware_b20130303) Process(packet packets.BanchoPacket) packets.BanchoPacket {
	switch packet.PacketId {
	case packets.BanchoProtocolNegotiation:
		return packet
	default:
		return packet
		panic("Unimplemented!")
	}
}

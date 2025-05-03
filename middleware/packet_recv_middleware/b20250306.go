package packet_recv_middleware

import "bancho-kill/packets"

type ReceiverMiddleware_b20250306 struct{}

func (ReceiverMiddleware_b20250306) Process(packet packets.BanchoPacket) packets.BanchoPacket {
	//this is currently the most recent client version with packet layout changes.
	//meaning the packets that arrive here are always in the very latest format.
	//which does mean it matches the cross version packet structures, so no conversion is necessary
	return packet
}

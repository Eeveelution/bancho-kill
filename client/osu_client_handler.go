package client

import (
	"bancho-kill/packets"
	"bytes"
	"encoding/binary"
)

func (client *OsuClient) HandleIncomingPacket(packet packets.BanchoPacket) {
	packetDataReader := bytes.NewBuffer(packet.PacketData)

	switch packet.PacketId {
	case packets.BanchoLoginReply:
		var userId int32

		binary.Read(packetDataReader, binary.LittleEndian, &userId)

		client.OwnUserData.UserID = uint64(userId)
	case packets.BanchoProtocolNegotiation:
		var protocolVer int32

		binary.Read(packetDataReader, binary.LittleEndian, &protocolVer)

		client.BanchoProtocolVersion = protocolVer
	case packets.BanchoHandleOsuUpdate:
		newStatsAndStatus := packets.Read[packets.UserStats](packetDataReader)

		if newStatsAndStatus.UserId == int32(client.OwnUserData.UserID) {
			convertedStruct := UserStats{
				UserID:      newStatsAndStatus.UserId,
				Mode:        newStatsAndStatus.Status.Playmode,
				Rank:        newStatsAndStatus.Rank,
				RankedScore: newStatsAndStatus.RankedScore,
				TotalScore:  newStatsAndStatus.TotalScore,
				Accuracy:    newStatsAndStatus.Accuracy,
				Playcount:   newStatsAndStatus.Playcount,
				Performance: newStatsAndStatus.Performance,
			}

			switch newStatsAndStatus.Status.Playmode {
			case packets.OsuGamemodeOsu:
				client.OwnUserData.StatsOsu = convertedStruct
			case packets.OsuGamemodeTaiko:
				client.OwnUserData.StatsTaiko = convertedStruct
			case packets.OsuGamemodeCatch:
				client.OwnUserData.StatsCatch = convertedStruct
			case packets.OsuGamemodeMania:
				client.OwnUserData.StatsMania = convertedStruct
			}
		}
	}
}

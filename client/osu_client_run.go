package client

import (
	"bancho-kill/packets"
	"bytes"
	"fmt"
)

func (client *OsuClient) MaintainClient() {

}

func (client *OsuClient) ReceiveData(data []byte) {
	readBuffer := bytes.NewBuffer(data)
	readIndex := 0

	for readIndex < len(data) {
		read, readPacket, failedRead := packets.ReadBanchoPacketHeader(readBuffer)

		readIndex += read

		if failedRead {
			continue
		}

		crossVersionPacket := client.PacketReaderMiddleware.Process(readPacket)

		client.PacketIncomingQueue <- crossVersionPacket

		fmt.Printf("%s\n", packets.GetPacketName(readPacket.PacketId))
	}
}

func (client *OsuClient) sendDataTcp(packet []byte) {

}

func (client *OsuClient) sendDataHttp(packet []byte) {

}

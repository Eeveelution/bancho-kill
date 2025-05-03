package client

import (
	"bancho-kill/packets"
	"bytes"
	"context"
	"time"
)

func (client *OsuClient) MaintainClient(ctx context.Context) {
	var httpPollTicker *time.Ticker

	if client.kind == ClientKindHttp {
		httpPollTicker = time.NewTicker(client.PacketPollRate * time.Millisecond)
	}

	for {
		select {
		case <-ctx.Done():
			if httpPollTicker != nil {
				httpPollTicker.Stop()
			}
		case <-httpPollTicker.C:

		}
	}
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
		client.PacketHistory = append(client.PacketHistory, crossVersionPacket)
	}
}

func (client *OsuClient) sendDataTcp(packet []byte) {

}

func (client *OsuClient) sendDataHttp(packet []byte) {

}

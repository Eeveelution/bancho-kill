package client

import (
	"bancho-kill/packets"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (client *OsuClient) maintainHttp(ctx context.Context) {
	version := fmt.Sprintf("b%d", client.clientVersion)

	httpPollTicker := time.NewTicker(client.PacketPollRate * time.Millisecond)
	httpClient := &http.Client{}

	for {
		select {
		case <-ctx.Done():
			if httpPollTicker != nil {
				httpPollTicker.Stop()
			}
		case <-httpPollTicker.C:
			if client.banchoToken == "" {
				continue
			}

			sendBuffer := new(bytes.Buffer)

			if len(client.PacketOutgoingQueue) == 0 {
				client.PacketOutgoingQueue <- packets.SendEmpty(packets.OsuPong)
			}

			for len(client.PacketOutgoingQueue) != 0 {
				currentPacket := <-client.PacketOutgoingQueue

				sendBuffer.Write(currentPacket)
			}

			req, err := http.NewRequest("POST", client.serverAddr, sendBuffer)

			if err != nil {
				continue
			}

			req.Header.Add("User-Agent", "osu!")
			req.Header.Add("osu-token", client.banchoToken)
			req.Header.Add("osu-version", version)

			resp, err := httpClient.Do(req)

			if err != nil {
				continue
			}

			body, err := io.ReadAll(resp.Body)

			if err != nil {
				continue
			}

			client.ReceiveData(body)
		case packet := <-client.PacketIncomingQueue:
			client.HandleIncomingPacket(packet)
		}
	}
}

func (client *OsuClient) tcpDataReceiver() {
	//make a 32kb Buffer to read stuff
	readBuffer := make([]byte, 32768)

	for client.continueRunning {
		read, readErr := client.conn.Read(readBuffer)

		if readErr != nil {
			//We don't clean up as we may not need to
			continue
		}

		//Get the bytes that were actually read
		packetData := readBuffer[:read]

		client.ReceiveData(packetData)
	}
}

func (client *OsuClient) maintainTcp(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			client.conn.Close()
			client.continueRunning = false
		case packet := <-client.PacketOutgoingQueue:
			client.conn.Write(packet)
		case packet := <-client.PacketIncomingQueue:
			client.HandleIncomingPacket(packet)
		}
	}
}

func (client *OsuClient) MaintainClient(ctx context.Context) {
	if client.kind == ClientKindHttp {
		client.maintainHttp(ctx)
	} else {
		go client.tcpDataReceiver()
		client.maintainTcp(ctx)
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

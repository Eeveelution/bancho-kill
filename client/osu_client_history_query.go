package client

import (
	"bancho-kill/packets"
	"fmt"
	"time"
)

func (client *OsuClient) hasReceivedPacket(packetId uint16) bool {
	if client == nil {
		return false
	}

	for _, packet := range client.PacketHistory {
		if packet.PacketId == packetId {
			return true
		}
	}

	return false
}

func (client *OsuClient) isTrueOnAboveVersion(version int, expression bool) bool {
	if client == nil {
		return false
	}

	if client.clientVersion < version {
		return true
	}

	return expression
}

func (client *OsuClient) isTrueOnBelowVersion(version int, expression bool) bool {
	if client == nil {
		return false
	}

	if client.clientVersion > version {
		return true
	}

	return expression
}

// ---------------------------------- PacketID Received functions

func (client *OsuClient) AssertPacketIdReceived(packetId uint16, desc string) *OsuClient {
	if client == nil {
		return nil
	}

	if client.hasReceivedPacket(packetId) {
		return client
	}

	client.TestFailFunc(fmt.Sprintf("%s expected, but never received. %s", packets.GetPacketName(packetId), desc))

	return nil
}

func (client *OsuClient) WarnIfPacketIdNotReceived(packetId uint16, desc string) *OsuClient {
	if client == nil {
		return nil
	}

	if client.hasReceivedPacket(packetId) {
		return client
	}

	client.TestFailFunc(fmt.Sprintf("%s expected, but never received. %s", packets.GetPacketName(packetId), desc))

	return client
}

func (client *OsuClient) AssertIfPacketNotReceivedAboveVersions(version int, packetId uint16, desc string) *OsuClient {
	if client == nil {
		return nil
	}

	if client.isTrueOnAboveVersion(version, client.hasReceivedPacket(packetId)) {
		return client
	}

	client.TestFailFunc(fmt.Sprintf("%s expected, but never received. %s", packets.GetPacketName(packetId), desc))

	return nil
}

func (client *OsuClient) WarnIfPacketNotReceivedAboveVersions(version int, packetId uint16, desc string) *OsuClient {
	if client == nil {
		return nil
	}

	if client.isTrueOnAboveVersion(version, client.hasReceivedPacket(packetId)) {
		return client
	}

	client.TestWarnFunc(fmt.Sprintf("%s expected, but never received. %s", packets.GetPacketName(packetId), desc))

	return client
}

// ---------------------------------- Waiting Functions

func (client *OsuClient) Wait(duration time.Duration) *OsuClient {
	if client == nil {
		return nil
	}

	time.Sleep(duration)

	return client
}

// ---------------------------------- General Functions

func (client *OsuClient) Assert(desc string, assertFunc func(OsuClient) bool) *OsuClient {
	if client == nil {
		return nil
	}

	if !assertFunc(*client) {
		client.TestFailFunc(desc)

		return nil
	}

	return client
}

func (client *OsuClient) WarnOn(desc string, assertFunc func(OsuClient) bool) *OsuClient {
	if client == nil {
		return nil
	}

	if !assertFunc(*client) {
		client.TestWarnFunc(desc)
	}

	return client
}

func (client *OsuClient) AssertOnVersionsAbove(version int, desc string, assertFunc func(OsuClient) bool) *OsuClient {
	if client == nil {
		return nil
	}

	if client.isTrueOnAboveVersion(version, assertFunc(*client)) {
		return client
	}

	client.TestFailFunc(desc)

	return nil
}

func (client *OsuClient) AssertOnVersionsBelow(version int, desc string, assertFunc func(OsuClient) bool) *OsuClient {
	if client == nil {
		return nil
	}

	if client.isTrueOnBelowVersion(version, assertFunc(*client)) {
		return client
	}

	client.TestFailFunc(desc)

	return nil
}

func (client *OsuClient) WarnOnVersionsAbove(version int, desc string, assertFunc func(OsuClient) bool) *OsuClient {
	if client == nil {
		return nil
	}

	if client.isTrueOnAboveVersion(version, assertFunc(*client)) {
		return client
	}

	client.TestWarnFunc(desc)

	return client
}

func (client *OsuClient) WarnOnVersionsBelow(version int, desc string, assertFunc func(OsuClient) bool) *OsuClient {
	if client == nil {
		return nil
	}

	if client.isTrueOnBelowVersion(version, assertFunc(*client)) {
		return client
	}

	client.TestWarnFunc(desc)

	return client
}

// ---------------------------------- Chat related checks

func (client *OsuClient) AssertJoinedChatChannel(channelName string) *OsuClient {
	if client == nil {
		return nil
	}

	for _, channel := range client.JoinedChannels {
		if channel.Name == channelName {
			return client
		}
	}

	client.TestFailFunc(fmt.Sprintf("Expected to be in #%s, never received a Join Success.", channelName))

	return nil
}

func (client *OsuClient) AssertKnowsOf(username string, desc string) *OsuClient {
	if client == nil {
		return nil
	}

	for _, users := range client.PresentUsers {
		if users.Username == username {
			return client
		}
	}

	client.TestFailFunc(fmt.Sprintf("Expected to have Presence information on %s, none received. %s", username, desc))

	return nil
}

func (client *OsuClient) WarnIfDoesntKnowOf(username string, desc string) *OsuClient {
	if client == nil {
		return nil
	}

	for _, users := range client.PresentUsers {
		if users.Username == username {
			return client
		}
	}

	client.TestWarnFunc(fmt.Sprintf("Expected to have Presence information on %s, none received. %s", username, desc))

	return client
}

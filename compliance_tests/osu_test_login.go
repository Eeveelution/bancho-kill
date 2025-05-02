package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/packets"
	"strings"
	"time"
)

func TestOsuLogin(ctx *TestContext) {
	ctx.CurrentTestNumber++

	createdClient, err := client.CreateOsuClient(ctx.OsuClientKind, ctx.OsuClientVersion, ctx.ServerAddress, ctx.Fail, ctx.Warn)

	if err != nil {
		ctx.FailWithErrorAndMessage("Failed to create client.", err)
	}

	createdClient.PacketReaderMiddleware = ctx.PacketRecvMiddleware

	createdClient.
		Login(ctx.LoginInformation).
		Wait(500*time.Millisecond).
		Assert("User has authenticated successfully as their own User.", func(oc client.OsuClient) bool {
			hasId := oc.OwnUserData.UserID >= 1
			hasExpectedUsername := strings.EqualFold(oc.OwnUserData.Username, ctx.LoginInformation.Username)

			return hasId && hasExpectedUsername
		}).
		Assert("User has received statistics for at least one mode.", func(oc client.OsuClient) bool {
			osu := oc.OwnUserData.HasStatsFor(packets.OsuGamemodeOsu)
			taiko := oc.OwnUserData.HasStatsFor(packets.OsuGamemodeTaiko)
			catch := oc.OwnUserData.HasStatsFor(packets.OsuGamemodeCatch)
			mania := oc.OwnUserData.HasStatsFor(packets.OsuGamemodeMania)

			return osu || taiko || catch || mania
		}).
		AssertKnowsOf(ctx.LoginInformation.Username, "We should see at the very least our own user.").
		AssertJoinedChatChannel("osu").
		WarnIfPacketIdNotReceivedOnVersionsAbove(425, packets.BanchoLoginPermissions, "Should have received a BanchoLoginPermissions above b425 to let the client know it's status.").
		WarnIfPacketIdNotReceivedOnVersionsAbove(535, packets.BanchoTitleUpdate, "osu!bancho does send a BanchoTitleUpdate on login, to make sure the clients Image is fully refreshed.")
}

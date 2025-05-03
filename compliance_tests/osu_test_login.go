package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/packets"
	"context"
	"strings"
	"time"
)

func TestSuccessfulOsuLogin(ctx *TestContext, testClient *client.OsuClient) {
	maintainCtx, cancelMaintain := context.WithCancel(context.Background())

	go testClient.MaintainClient(maintainCtx)

	testClient.
		Login(ctx.LoginInformation).
		Wait(500*time.Millisecond).
		Assert("User should have authenticated successfully as their own User.", func(oc client.OsuClient) bool {
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
		WarnIfPacketNotReceivedAboveVersions(425, packets.BanchoLoginPermissions, "Should have received a BanchoLoginPermissions above b425 to let the client know it's status.").
		WarnIfPacketNotReceivedAboveVersions(535, packets.BanchoTitleUpdate, "osu!bancho does send a BanchoTitleUpdate on login, to make sure the clients Image is fully refreshed.")

	time.Sleep(time.Second)

	cancelMaintain()
}

func TestLoginFailure(ctx *TestContext, testClient *client.OsuClient) {
	copiedLoginInfo := ctx.LoginInformation
	copiedLoginInfo.Password += "this is a incorrect password!"

	testClient.
		Login(copiedLoginInfo).
		Wait(500*time.Millisecond).
		Assert("Login should have failed.", func(oc client.OsuClient) bool {
			return oc.OwnUserData.UserID <= 0
		})
}

package compliance_tests

import "bancho-kill/client"

func TestOsuLogin(ctx *TestContext) {
	ctx.CurrentTestNumber++

	client, err := client.CreateOsuClient(ctx.OsuClientKind, ctx.ServerAddress)

	if err != nil {
		ctx.FailWithErrorAndMessage("Failed to create client.", err)
	}

	client.PacketMiddleware = append(client.PacketMiddleware, ctx.PacketMiddleware)

	client.Login("Furball", "ssh")
}

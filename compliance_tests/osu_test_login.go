package compliance_tests

import "bancho-kill/client"

func TestOsuLogin(ctx *TestContext) {
	ctx.CurrentTestNumber++

	client, err := client.CreateOsuClient(ctx.OsuClientKind, ctx.OsuClientVersion, ctx.ServerAddress)

	if err != nil {
		ctx.FailWithErrorAndMessage("Failed to create client.", err)
	}

	client.PacketReaderMiddleware = ctx.PacketRecvMiddleware

	client.Login(ctx.Username, ctx.Password, 8, false, false, "b4dcc68511aa60e190541af7668e442f:runningunderwine:b4ec3c4334a0249dae95c284ec5983df:4cfdc2e157eefe6facb983b1d557b3a1:4cfdc2e157eefe6facb983b1d557b3a1")
}

package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware"
	"bancho-kill/packets"
)

type TestContext struct {
	CurrentTestNumber    int
	TotalFails           int
	TotalWarnings        int
	ServerAddress        string
	OsuClientKind        client.ClientKind
	OsuClientVersion     int
	PacketRecvMiddleware middleware.PacketMiddleware

	LoginInformation packets.InitialLoginInformation
}

func (ctx TestContext) Fail(err string) {

}

func (ctx TestContext) FailWithError(err error) {

}

func (ctx TestContext) FailWithErrorAndMessage(str string, err error) {

}

func (ctx TestContext) Warn(warn string) {

}

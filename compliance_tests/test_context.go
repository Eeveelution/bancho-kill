package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware"
)

type TestContext struct {
	CurrentTestNumber int
	ServerAddress     string
	OsuClientKind     client.ClientKind
	PacketMiddleware  middleware.Middleware
}

func (ctx TestContext) Fail(err string) {

}

func (ctx TestContext) FailWithError(err error) {

}

func (ctx TestContext) FailWithErrorAndMessage(str string, err error) {

}

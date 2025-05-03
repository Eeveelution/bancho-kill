package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware"
	"bancho-kill/packets"
)

type WarningOrError struct {
	IsError     bool
	Error       error
	Description string
}

type TestContext struct {
	CurrentTestNumber    int
	CurrentTestName      string
	TotalFails           int
	TotalWarnings        int
	ServerAddress        string
	OsuClientKind        client.ClientKind
	OsuClientVersion     int
	PacketRecvMiddleware middleware.PacketMiddleware

	LoginInformation packets.InitialLoginInformation

	Warnings map[string][]WarningOrError
}

func (ctx *TestContext) Fail(err string) {
	ctx.Warnings[ctx.CurrentTestName] = append(ctx.Warnings[ctx.CurrentTestName], WarningOrError{
		IsError:     true,
		Error:       nil,
		Description: err,
	})
}

func (ctx *TestContext) FailWithError(err error) {
	ctx.Warnings[ctx.CurrentTestName] = append(ctx.Warnings[ctx.CurrentTestName], WarningOrError{
		IsError:     true,
		Error:       err,
		Description: "",
	})
}

func (ctx *TestContext) FailWithErrorAndMessage(str string, err error) {
	ctx.Warnings[ctx.CurrentTestName] = append(ctx.Warnings[ctx.CurrentTestName], WarningOrError{
		IsError:     true,
		Error:       err,
		Description: str,
	})
}

func (ctx *TestContext) Warn(warn string) {
	ctx.Warnings[ctx.CurrentTestName] = append(ctx.Warnings[ctx.CurrentTestName], WarningOrError{
		IsError:     false,
		Error:       nil,
		Description: "Warn",
	})
}

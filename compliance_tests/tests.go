package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware/packet_recv_middleware"
	"fmt"
	"os"
	"strings"
)

var complianceTests map[string]func(*TestContext) = map[string]func(*TestContext){
	"1 - Testing a basic osu! Login": TestOsuLogin,
}

func RunAllComplianceTests(addr string) {
	complianceLogin, err := os.ReadFile("compliance_login.txt")

	if err != nil {
		fmt.Printf("Failed to read login details for the compliance test account!")

		return
	}

	complianceLoginAsStr := string(complianceLogin)
	complianceLoginSplit := strings.Split(complianceLoginAsStr, "\n")

	username := complianceLoginSplit[0]
	password := complianceLoginSplit[1]

	context := TestContext{
		CurrentTestNumber:    0,
		ServerAddress:        addr,
		OsuClientKind:        client.ClientKindHttp,
		OsuClientVersion:     20130303,
		Username:             username,
		Password:             password,
		PacketRecvMiddleware: packet_recv_middleware.ReceiverMiddleware_b20130303{},
	}

	TestOsuLogin(&context)
}

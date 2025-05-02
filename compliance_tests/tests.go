package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware/packet_recv_middleware"
	"bancho-kill/packets"
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
		CurrentTestNumber: 0,
		ServerAddress:     addr,
		OsuClientKind:     client.ClientKindHttp,
		OsuClientVersion:  20130303,
		LoginInformation: packets.InitialLoginInformation{
			Username:       username,
			Password:       password,
			Timezone:       8,
			ShowCity:       false,
			BlockFriendPms: false,
			ClientHash:     "b4dcc68511aa60e190541af7668e442f", //TODO: insert a actual client hash
			Adapter:        "runningunderwine",
			AdapterHash:    "b4ec3c4334a0249dae95c284ec5983df", //TODO: figure out how to do this correctly
			UninstallId:    "4cfdc2e157eefe6facb983b1d557b3a1", //TODO: figure out how to make these not dummies
			DiskSignature:  "4cfdc2e157eefe6facb983b1d557b3a1",
		},
		PacketRecvMiddleware: packet_recv_middleware.ReceiverMiddleware_b20130303{},
	}

	TestOsuLogin(&context)
}

package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware/packet_recv_middleware"
	"bancho-kill/packets"
	"fmt"
	"os"
	"strings"
	"time"
)

var complianceTests map[string]func(*TestContext, *client.OsuClient) = map[string]func(*TestContext, *client.OsuClient){
	"1.0 - Testing a basic osu! Login":     TestSuccessfulOsuLogin,
	"1.1 - Testing a incorrect osu! Login": TestLoginFailure,
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
		TotalFails:           0,
		TotalWarnings:        0,
		Warnings:             map[string][]WarningOrError{},
	}

	createdClient, err := client.CreateOsuClient(context.OsuClientKind, context.OsuClientVersion, context.ServerAddress, context.Fail, context.Warn)

	if err != nil {
		context.FailWithErrorAndMessage("Failed to create client.", err)
	}

	createdClient.PacketReaderMiddleware = context.PacketRecvMiddleware

	fmt.Print("\033[s")

	for testName, test := range complianceTests {
		context.CurrentTestNumber++
		context.CurrentTestName = testName

		fmt.Printf("[ RUNNING ] %s", testName)

		testStart := time.Now()

		test(&context, createdClient)

		elapsed := time.Since(testStart)

		warningsForCurrentTest := context.Warnings[testName]

		if len(warningsForCurrentTest) == 0 {
			fmt.Printf("\033[2K\r[ PASS ] (%dms) %s\n", elapsed.Milliseconds(), testName)
		} else {
			hasError := false

			for _, warning := range warningsForCurrentTest {
				if warning.IsError {
					hasError = true
					break
				}
			}

			if hasError {
				fmt.Printf("\033[2K\r[ FAIL ] (%dms) %s\n", elapsed.Milliseconds(), testName)
			} else {
				fmt.Printf("\033[2K\r[ WARN ] (%dms) %s\n", elapsed.Milliseconds(), testName)
			}

			for _, warning := range warningsForCurrentTest {
				if warning.IsError {
					fmt.Printf("- !!!! - ")
				} else {
					fmt.Printf("- ~~~~ - ")
				}

				if warning.Error != nil {
					fmt.Printf("%s", warning.Error.Error())
				}

				if warning.Description != "" {
					fmt.Printf("%s", warning.Description)
				}

				fmt.Printf("\n")
			}
		}
	}
}

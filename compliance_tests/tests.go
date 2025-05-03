package compliance_tests

import (
	"bancho-kill/client"
	"bancho-kill/middleware/packet_recv_middleware"
	"bancho-kill/packets"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/muesli/termenv"
)

type TestFunction func(*TestContext, *client.OsuClient)

type Test struct {
	TestName     string
	TestFunction TestFunction
}

type TestGrouping struct {
	GroupName string
	Tests     []Test
}

var complianceTests []TestGrouping = []TestGrouping{
	{
		GroupName: "osu! Login Tests",
		Tests: []Test{
			{
				TestName:     "Basic login until a \"Welcome to Bancho!\" should appear on the client.",
				TestFunction: TestSuccessfulOsuLogin,
			},
			{
				TestName:     "Testing a invalid login.",
				TestFunction: TestLoginFailure,
			},
		},
	},
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
		OsuClientVersion:  20250306,
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
		PacketRecvMiddleware: packet_recv_middleware.ReceiverMiddleware_b20250306{},
		TotalFails:           0,
		TotalWarnings:        0,
		Warnings:             map[string][]WarningOrError{},
	}

	createdClient, err := client.CreateOsuClient(context.OsuClientKind, context.OsuClientVersion, context.ServerAddress, context.Fail, context.Warn)

	if err != nil {
		context.FailWithErrorAndMessage("Failed to create client.", err)
	}

	createdClient.PacketReaderMiddleware = context.PacketRecvMiddleware

	output := termenv.NewOutput(os.Stdout)

	for groupIndex, group := range complianceTests {
		fmt.Printf("======== Beginning Test Group: %s\n", group.GroupName)

		for testIndex, test := range group.Tests {
			testIndex := fmt.Sprintf("%d.%d", groupIndex+1, testIndex+1)
			indiciesAndtestName := fmt.Sprintf("%s - %s", testIndex, test.TestName)

			context.CurrentTestName = testIndex

			fmt.Printf("[ RUNNING ] %s - %s", testIndex, test.TestName)

			testStart := time.Now()

			test.TestFunction(&context, createdClient)

			elapsed := time.Since(testStart)

			warningsForCurrentTest := context.Warnings[testIndex]

			if len(warningsForCurrentTest) == 0 {
				output.ClearLine()

				fmt.Printf("\r[ %s ] (%dms) %s\n", termenv.String("PASS").Foreground(termenv.ANSIGreen), elapsed.Milliseconds(), indiciesAndtestName)
			} else {
				output.ClearLine()

				hasError := false

				for _, warning := range warningsForCurrentTest {
					if warning.IsError {
						hasError = true
						break
					}
				}

				if hasError {
					fmt.Printf("\r[ %s ] (%dms) %s\n", termenv.String("FAIL").Foreground(termenv.ANSIRed), elapsed.Milliseconds(), indiciesAndtestName)
				} else {
					fmt.Printf("\r[ %s ] (%dms) %s\n", termenv.String("WARN").Foreground(termenv.ANSIYellow), elapsed.Milliseconds(), indiciesAndtestName)
				}

				for _, warning := range warningsForCurrentTest {
					line := ""

					if warning.IsError {
						line += "- !!!! - "
					} else {
						line += "- ~~~~ - "
					}

					if warning.Error != nil {
						line += warning.Error.Error() + " "
					}

					if warning.Description != "" {
						line += warning.Description
					}

					line += "\n"

					if warning.IsError {
						fmt.Printf("%s", termenv.String(line).Foreground(termenv.ANSIRed))
					} else {
						fmt.Printf("%s", termenv.String(line).Foreground(termenv.ANSIYellow))
					}
				}
			}
		}
	}
}

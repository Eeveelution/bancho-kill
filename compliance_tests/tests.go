package compliance_tests

import (
	"bancho-kill/client"
	"fmt"
	"os"
	"strings"
)

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
		Username:          username,
		Password:          password,
	}

	TestOsuLogin(&context)
}

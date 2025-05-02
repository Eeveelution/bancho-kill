package client

import (
	"bancho-kill/packets"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func (client *OsuClient) Login(loginInfo packets.InitialLoginInformation) *OsuClient {
	version := fmt.Sprintf("b%d", client.clientVersion)

	displayCityAsInt := "0"

	if loginInfo.ShowCity {
		displayCityAsInt = "1"
	}

	passwordHashed := md5.Sum([]byte(loginInfo.Password))
	passwordHashedString := hex.EncodeToString(passwordHashed[:])

	loginStr := fmt.Sprintf("%s\n%s\n%s|%d|%s|%s", loginInfo.Username, passwordHashedString, version, loginInfo.Timezone, displayCityAsInt, loginInfo.GetLoginClientHash(client.clientVersion))

	if client.clientVersion >= 20130303 {
		friendPmsAsInt := "0"

		if loginInfo.BlockFriendPms {
			friendPmsAsInt = "1"
		}

		loginStr += fmt.Sprintf("|%s\n", friendPmsAsInt)
	} else {
		loginStr += "\n"
	}

	if client.kind == ClientKindTcp {

	} else {
		httpClient := &http.Client{}

		req, err := http.NewRequest("POST", client.serverAddr, bytes.NewBufferString(loginStr))

		if err != nil {
			return nil
		}

		req.Header.Add("User-Agent", "osu!")
		req.Header.Add("osu-version", version)

		resp, err := httpClient.Do(req)

		if err != nil {
			return nil
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil
		}

		client.banchoToken = resp.Header.Get("cho-token")
		client.ReceiveData(body)
	}

	return client
}

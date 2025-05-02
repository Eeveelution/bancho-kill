package client

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

func (client *OsuClient) Login(username string, password string, timezone int, displayCity bool, blockingNonFriendPms bool, clientHash string) *OsuClient {
	version := fmt.Sprintf("b%d", client.ClientVersion)

	displayCityAsInt := "0"

	if displayCity {
		displayCityAsInt = "1"
	}

	passwordHashed := md5.Sum([]byte(password))
	passwordHashedString := hex.EncodeToString(passwordHashed[:])

	loginStr := fmt.Sprintf("%s\n%s\n%s|%d|%s|%s", username, passwordHashedString, version, timezone, displayCityAsInt, clientHash)

	if client.ClientVersion >= 20130303 {
		friendPmsAsInt := "0"

		if blockingNonFriendPms {
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

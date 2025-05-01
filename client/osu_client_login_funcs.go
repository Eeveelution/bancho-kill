package client

import "fmt"

func (client *OsuClient) Login(username string, password string, version string) {
	loginStr := fmt.Sprintf(`%s\n%s\n`)
}

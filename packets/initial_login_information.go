package packets

type InitialLoginInformation struct {
	Username       string
	Password       string
	Timezone       int8
	ShowCity       bool
	BlockFriendPms bool

	ClientHash    string
	Adapter       string
	AdapterHash   string
	UninstallId   string
	DiskSignature string
}

func (info InitialLoginInformation) GetLoginClientHash(version int) string {
	returnString := info.ClientHash + ":" + info.Adapter + ":" + info.AdapterHash

	//TODO figure out when uninstallid and disksignature got added

	return returnString
}

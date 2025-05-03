package packets

//The cross version packet structures are always the very latest client's packet structures.

type StatusUpdate struct {
	Action          uint8
	StatusText      string
	BeatmapChecksum string
	Mods            uint32
	Playmode        OsuGameMode
	BeatmapId       int32
}

type UserStats struct {
	UserId      int32
	Status      StatusUpdate
	RankedScore uint64
	Accuracy    float32
	Playcount   uint32
	TotalScore  uint64
	Rank        uint32
	Performance uint16
}

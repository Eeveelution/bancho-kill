package client

import "bancho-kill/packets"

type UserData struct {
	UserID        uint64
	Username      string
	Country       uint16
	Banned        bool
	Privileges    int32
	SilencedUntil uint64

	StatsOsu   UserStats
	StatsTaiko UserStats
	StatsCatch UserStats
	StatsMania UserStats

	Friends []FriendEntry
}

func (data *UserData) HasStatsFor(mode packets.OsuGameMode) bool {
	switch mode {
	case packets.OsuGamemodeOsu:
		return data.StatsOsu.UserID != 0
	case packets.OsuGamemodeTaiko:
		return data.StatsOsu.UserID != 0
	case packets.OsuGamemodeCatch:
		return data.StatsOsu.UserID != 0
	case packets.OsuGamemodeMania:
		return data.StatsOsu.UserID != 0
	}

	return false
}

type ChatChannel struct {
	Name      string
	Topic     string
	UserCount uint32
}

type FriendEntry struct {
	User1 uint64
	User2 uint64
}

type UserStats struct {
	UserID         int32
	Mode           packets.OsuGameMode
	Rank           uint32
	RankedScore    uint64
	TotalScore     uint64
	Performance    uint16
	Level          float64
	Accuracy       float32
	Playcount      uint32
	CountSSH       uint64
	CountSS        uint64
	CountSH        uint64
	CountS         uint64
	CountA         uint64
	CountB         uint64
	CountC         uint64
	CountD         uint64
	Hit300         uint64
	Hit100         uint64
	Hit50          uint64
	HitMiss        uint64
	HitGeki        uint64
	HitKatu        uint64
	ReplaysWatched uint64
	Playtime       uint64
}

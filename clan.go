package clash

// Clan represent a model that is fetched from API.
type Clan struct {
	BadgeUrls struct {
		Large  string `json:"large"`
		Medium string `json:"medium"`
		Small  string `json:"small"`
	} `json:"badgeUrls"`
	ClanLevel        int64  `json:"clanLevel"`
	ClanPoints       int64  `json:"clanPoints"`
	ClanVersusPoints int64  `json:"clanVersusPoints"`
	Description      string `json:"description"`
	IsWarLogPublic   bool   `json:"isWarLogPublic"`
	Location struct {
		CountryCode string `json:"countryCode"`
		ID          int64  `json:"id"`
		IsCountry   bool   `json:"isCountry"`
		Name        string `json:"name"`
	} `json:"location"`
	MemberList       []*Player `json:"memberList"`
	Members          int64     `json:"members"`
	Name             string    `json:"name"`
	RequiredTrophies int64     `json:"requiredTrophies"`
	Tag              string    `json:"tag"`
	Type             string    `json:"type"`
	WarFrequency     string    `json:"warFrequency"`
	WarLosses        int64     `json:"warLosses"`
	WarTies          int64     `json:"warTies"`
	WarWinStreak     int64     `json:"warWinStreak"`
	WarWins          int64     `json:"warWins"`
}

package clash

// Player model from the API
type Player struct {
	Achievements []struct {
		CompletionInfo string `json:"completionInfo"`
		Info           string `json:"info"`
		Name           string `json:"name"`
		Stars          int64  `json:"stars"`
		Target         int64  `json:"target"`
		Value          int64  `json:"value"`
		Village        string `json:"village"`
	} `json:"achievements"`
	AttackWins         int64 `json:"attackWins"`
	BestTrophies       int64 `json:"bestTrophies"`
	BestVersusTrophies int64 `json:"bestVersusTrophies"`
	BuilderHallLevel   int64 `json:"builderHallLevel"`
	Clan struct {
		BadgeUrls struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Small  string `json:"small"`
		} `json:"badgeUrls"`
		ClanLevel int64  `json:"clanLevel"`
		Name      string `json:"name"`
		Tag       string `json:"tag"`
	} `json:"clan"`
	DefenseWins       int64 `json:"defenseWins"`
	Donations         int64 `json:"donations"`
	DonationsReceived int64 `json:"donationsReceived"`
	ExpLevel          int64 `json:"expLevel"`
	Heroes []struct {
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Name     string `json:"name"`
		Village  string `json:"village"`
	} `json:"heroes"`
	League struct {
		IconUrls struct {
			Medium string `json:"medium"`
			Small  string `json:"small"`
			Tiny   string `json:"tiny"`
		} `json:"iconUrls"`
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"league"`
	LegendStatistics struct {
		BestSeason struct {
			ID       string `json:"id"`
			Rank     int64  `json:"rank"`
			Trophies int64  `json:"trophies"`
		} `json:"bestSeason"`
		CurrentSeason struct {
			Trophies int64 `json:"trophies"`
		} `json:"currentSeason"`
		LegendTrophies int64 `json:"legendTrophies"`
	} `json:"legendStatistics"`
	Name string `json:"name"`
	Role string `json:"role"`
	Spells []struct {
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Name     string `json:"name"`
		Village  string `json:"village"`
	} `json:"spells"`
	Tag           string `json:"tag"`
	TownHallLevel int64  `json:"townHallLevel"`
	Troops []struct {
		Level    int64  `json:"level"`
		MaxLevel int64  `json:"maxLevel"`
		Name     string `json:"name"`
		Village  string `json:"village"`
	} `json:"troops"`
	Trophies             int64 `json:"trophies"`
	VersusBattleWinCount int64 `json:"versusBattleWinCount"`
	VersusBattleWins     int64 `json:"versusBattleWins"`
	VersusTrophies       int64 `json:"versusTrophies"`
	WarStars             int64 `json:"warStars"`
}

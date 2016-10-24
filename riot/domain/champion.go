package domain

type Champion struct {
	Active            bool
	BotEnabled        bool
	BotMmEnabled      bool
	FreeToPlay        bool
	Id                int64
	RankedPlayEnabled bool
}

type ChampionStatic struct {
	Name string
}

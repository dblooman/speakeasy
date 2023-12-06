package openapi

type NatureModel struct {
	DecreasedStat              interface{} `json:"decreased_stat"`
	HatesFlavor                interface{} `json:"hates_flavor"`
	ID                         int         `json:"id"`
	IncreasedStat              interface{} `json:"increased_stat"`
	LikesFlavor                interface{} `json:"likes_flavor"`
	MoveBattleStylePreferences []struct {
		HighHpPreference int `json:"high_hp_preference"`
		LowHpPreference  int `json:"low_hp_preference"`
		MoveBattleStyle  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move_battle_style"`
	} `json:"move_battle_style_preferences"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokeathlonStatChanges []struct {
		MaxChange      int `json:"max_change"`
		PokeathlonStat struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokeathlon_stat"`
	} `json:"pokeathlon_stat_changes"`
}

type StatModel struct {
	AffectingMoves struct {
		Decrease []interface{} `json:"decrease"`
		Increase []interface{} `json:"increase"`
	} `json:"affecting_moves"`
	AffectingNatures struct {
		Decrease []interface{} `json:"decrease"`
		Increase []interface{} `json:"increase"`
	} `json:"affecting_natures"`
	Characteristics []struct {
		URL string `json:"url"`
	} `json:"characteristics"`
	GameIndex       int         `json:"game_index"`
	ID              int         `json:"id"`
	IsBattleOnly    bool        `json:"is_battle_only"`
	MoveDamageClass interface{} `json:"move_damage_class"`
	Name            string      `json:"name"`
	Names           []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
}

type PokemonModel struct {
	ID                     int         `json:"id"`
	Name                   string      `json:"name"`
	BaseExperience         int         `json:"base_experience"`
	Height                 int         `json:"height"`
	IsDefault              bool        `json:"is_default"`
	Order                  int         `json:"order"`
	Weight                 int         `json:"weight"`
	Abilities              []Ability   `json:"abilities"`
	Forms                  []Form      `json:"forms"`
	GameIndices            []GameIndex `json:"game_indices"`
	HeldItems              []HeldItem  `json:"held_items"`
	LocationAreaEncounters string      `json:"location_area_encounters"`
	Moves                  []Move      `json:"moves"`
	Species                Species     `json:"species"`
	Sprites                Sprites     `json:"sprites"`
	Stats                  []Stat      `json:"stats"`
	Types                  []Type      `json:"types"`
	PastTypes              []PastType  `json:"past_types"`
}

type Ability struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
	Ability  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"ability"`
}

type Form struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GameIndex struct {
	GameIndex int `json:"game_index"`
	Version   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
}

type HeldItem struct {
	Item struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type VersionDetail struct {
	Rarity  int `json:"rarity"`
	Version struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version"`
}

type Move struct {
	Move struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}

type VersionGroupDetail struct {
	LevelLearnedAt int `json:"level_learned_at"`
	VersionGroup   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"version_group"`
	MoveLearnMethod struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_learn_method"`
}

type Species struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	BackDefault      string       `json:"back_default"`
	BackFemale       interface{}  `json:"back_female"`
	BackShiny        string       `json:"back_shiny"`
	BackShinyFemale  interface{}  `json:"back_shiny_female"`
	FrontDefault     string       `json:"front_default"`
	FrontFemale      interface{}  `json:"front_female"`
	FrontShiny       string       `json:"front_shiny"`
	FrontShinyFemale interface{}  `json:"front_shiny_female"`
	Other            OtherSprites `json:"other"`
	Versions         Versions     `json:"versions"`
}

type OtherSprites struct {
	DreamWorld      DreamWorldSprites      `json:"dream_world"`
	Home            HomeSprites            `json:"home"`
	OfficialArtwork OfficialArtworkSprites `json:"official-artwork"`
}

type DreamWorldSprites struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}

type HomeSprites struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type OfficialArtworkSprites struct {
	FrontDefault string `json:"front_default"`
}

type Versions struct {
	Generation Generation `json:"generation"`
}

type RedBlueSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

type YellowSprites struct {
	BackDefault  string `json:"back_default"`
	BackGray     string `json:"back_gray"`
	FrontDefault string `json:"front_default"`
	FrontGray    string `json:"front_gray"`
}

type Generation struct {
	Crystal                CrystalSprites                `json:"crystal"`
	Gold                   GoldSprites                   `json:"gold"`
	Silver                 SilverSprites                 `json:"silver"`
	RedBlue                RedBlueSprites                `json:"red-blue"`
	Yellow                 YellowSprites                 `json:"yellow"`
	Emerald                EmeraldSprites                `json:"emerald"`
	FireredLeafgreen       FireredLeafgreenSprites       `json:"firered-leafgreen"`
	RubySapphire           RubySapphireSprites           `json:"ruby-sapphire"`
	DiamondPearl           DiamondPearlSprites           `json:"diamond-pearl"`
	HeartgoldSoulsilver    HeartgoldSoulsilverSprites    `json:"heartgold-soulsilver"`
	Platinum               PlatinumSprites               `json:"platinum"`
	BlackWhite             BlackWhiteSprites             `json:"black-white"`
	OmegarubyAlphasapphire OmegarubyAlphasapphireSprites `json:"omegaruby-alphasapphire"`
	XY                     XYSprites                     `json:"x-y"`
	Icons                  IconsSprites                  `json:"icons"`
	UltraSunUltraMoon      UltraSunUltraMoonSprites      `json:"ultra-sun-ultra-moon"`
}

type CrystalSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type GoldSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type SilverSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type EmeraldSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type FireredLeafgreenSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type RubySapphireSprites struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type DiamondPearlSprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type HeartgoldSoulsilverSprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type PlatinumSprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type BlackWhiteSprites struct {
	Animated         AnimatedSprites `json:"animated"`
	BackDefault      string          `json:"back_default"`
	BackFemale       interface{}     `json:"back_female"`
	BackShiny        string          `json:"back_shiny"`
	BackShinyFemale  interface{}     `json:"back_shiny_female"`
	FrontDefault     string          `json:"front_default"`
	FrontFemale      interface{}     `json:"front_female"`
	FrontShiny       string          `json:"front_shiny"`
	FrontShinyFemale interface{}     `json:"front_shiny_female"`
}

type AnimatedSprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type OmegarubyAlphasapphireSprites struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type XYSprites struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type IconsSprites struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}

type UltraSunUltraMoonSprites struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type PastType struct {
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	Types []Type `json:"types"`
}

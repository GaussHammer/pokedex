package internal

type Pokemon struct {
	Name            string  `json:"name"`
	Base_experience int     `json:"base_experience"`
	Height          int     `json:"height"`
	Weight          int     `json:"weight"`
	Stats           []Stats `json:"stats"`
	Types           []Types `json:"types"`
}

type Stats struct {
	Stat struct {
		Name string `json:"name"`
	}
	Base_stat int `json:"base_stat"`
}

type Types struct {
	Type struct {
		Name string `json:"name"`
	}
}

var Pokedex map[string]Pokemon

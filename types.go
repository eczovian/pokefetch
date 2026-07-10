package main

type PokemonGeneralData struct {
	Name string `json:"name"`
}

type PokemonTypeData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonTypeRequestData struct {
	Name           string              `json:"name"`
	LocalizedNames []PokemonLocaleData `json:"names"`
}

type PokemonLocaleData struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

type Language struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonType struct {
	TypeData PokemonGeneralData `json:"type"`
}

type PokemonData struct {
	Types  []PokemonType `json:"types"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Id     int           `json:"id"`
}

type PokemonSpeciesGenera struct {
	Genus    string             `json:"genus"`
	Language PokemonGeneralData `json:"language"`
}

type PokemonSpeciesName struct {
	Name     string             `json:"name"`
	Language PokemonGeneralData `json:"language"`
}

type PokemonSpeciesFlavorText struct {
	FlavorText string             `json:"flavor_text"`
	Language   PokemonGeneralData `json:"language"`
}

type PokemonSpeciesData struct {
	Names             []PokemonSpeciesName       `json:"names"`
	Genera            []PokemonSpeciesGenera     `json:"genera"`
	FlavorTextEntries []PokemonSpeciesFlavorText `json:"flavor_text_entries"`
}

type LocaleData struct {
	Names []PokemonLocaleData `json:"names"`
}

type OuterLocaleData struct {
	Height LocaleData `json:"height"`
	Weight LocaleData `json:"weight"`
}

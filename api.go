package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
  _ "embed"
)

const (
	PokemonAPI        = "https://pokeapi.co/api/v2/pokemon"
	PokemonSpeciesAPI = "https://pokeapi.co/api/v2/pokemon-species"
	PokemonTypeAPI    = "https://pokeapi.co/api/v2/type"
)

//go:embed missingno
var missingNoSprite []byte

func fetchPokemonData(pokemon string) PokemonData {
	return fetchData[PokemonData](fmt.Sprintf("%s/%s", PokemonAPI, pokemon))
}

func fetchPokemonSpeciesData(pokemon string) PokemonSpeciesData {
	return fetchData[PokemonSpeciesData](fmt.Sprintf("%s/%s", PokemonSpeciesAPI, pokemon))
}

func fetchPokemonTypeData(pokeType string) PokemonTypeRequestData {
	return fetchData[PokemonTypeRequestData](fmt.Sprintf("%s/%s", PokemonTypeAPI, pokeType))
}

func isValidPokemonName(name string) bool {
	resp, err := http.Get(fmt.Sprintf("%s/%s", PokemonAPI, name))

	if err != nil {
		return false
	}

	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func fetchData[T any](url string) T {
  var dataReader io.Reader
	byteData, err := getCachedFileContentOrError(url)
  if err != nil {
    resp, err := http.Get(url)
    if err != nil {
      log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err:= io.ReadAll(resp.Body)
    writeRequestBytesToFile(body,url)
    dataReader = bytes.NewReader(body)
  } else {
    dataReader = bytes.NewReader(byteData)
  }

	var data T
	if err := json.NewDecoder(dataReader).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return data
}

func fetchPokemonImage(url string) string {
	var readableData []byte
	data, err := getCachedFileContentOrError(url)
	if err != nil {
		resp, err := http.Get(url)
    if resp.StatusCode!=200 {
      return string(missingNoSprite)
    }
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
    writeRequestBytesToFile(body, url)
		readableData = body
	} else {
		readableData = data
	}

	return string(readableData)
}

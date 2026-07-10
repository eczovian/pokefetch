package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed locale.json
var localeText []byte;

func parseDataLocaleWeightAndHeightData() OuterLocaleData{
  var data OuterLocaleData;
  if err:= json.NewDecoder(bytes.NewReader(localeText)).Decode(&data); err!=nil {
    log.Fatal(err)
  }
  return data;
}


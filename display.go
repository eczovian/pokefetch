package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func getEnglishName(names []PokemonSpeciesName) string {
  return getLocalizedName(names, "en")
}

func getLocalizedHeight(localeData OuterLocaleData, locale string) string{
  for _, name := range localeData.Height.Names {
    if name.Language.Name == locale {
      return name.Name
    }
  }
  return ""
}

func getLocalizedWeight(localeData OuterLocaleData, locale string) string{
  for _, name := range localeData.Weight.Names {
    if name.Language.Name == locale {
      return name.Name
    }
  }
  return ""
}

func getLocalizedName(names []PokemonSpeciesName, locale string) string {
	for _, name := range names {
		if name.Language.Name == locale {
			return name.Name
		}
	}
	return ""
}

func getLocalizedGenus(genera []PokemonSpeciesGenera, locale string) string {
	for _, genus := range genera {
		if genus.Language.Name == locale {
			return genus.Genus
		}
	}
	return ""
}

func getLocalizedFlavorText(entries []PokemonSpeciesFlavorText, locale string) string {
	for i := len(entries) - 1; i >= 0; i-- {
		if entries[i].Language.Name == locale {
			return strings.ReplaceAll(entries[i].FlavorText, "\n", " ")
		}
	}
	return ""
}

func getLocalizedTypeBadges(entries []PokemonTypeRequestData, locale string) string{
  var badges []string;
  for i:= len(entries) -1; i >= 0; i--{
    for j:= len(entries[i].LocalizedNames) -1; j >=0; j--{
      if(entries[i].LocalizedNames[j].Language.Name == locale) {
        color := pokemonTypeColor(entries[i].Name)
        badges = append(badges, createTextBadge(entries[i].LocalizedNames[j].Name,color,false))
      }
    }
  }
  return strings.Join(badges, " ")
}

func getTypeBadges(types []PokemonType) string {
	var badges []string
	for _, t := range types {
		color := pokemonTypeColor(t.TypeData.Name)
		badges = append(badges, createTextBadge(strings.ToUpper(t.TypeData.Name), color, false))
	}
	return strings.Join(badges, " ")
}

func createTextBadge(text string, color lipgloss.Color, bold bool) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(color).Padding(0, 1)
	if bold {
		style = style.Bold(true)
	}
	return style.Render(text)
}

func pokemonTypeColor(pokemonType string) lipgloss.Color {
	colors := map[string]lipgloss.Color{
		"normal":   lipgloss.Color("15"),
		"fire":     lipgloss.Color("9"),
		"water":    lipgloss.Color("12"),
		"electric": lipgloss.Color("11"),
		"grass":    lipgloss.Color("10"),
		"ice":      lipgloss.Color("14"),
		"fighting": lipgloss.Color("1"),
		"poison":   lipgloss.Color("5"),
		"ground":   lipgloss.Color("3"),
		"flying":   lipgloss.Color("13"),
		"psychic":  lipgloss.Color("13"),
		"bug":      lipgloss.Color("2"),
		"rock":     lipgloss.Color("3"),
		"ghost":    lipgloss.Color("5"),
		"dragon":   lipgloss.Color("13"),
		"dark":     lipgloss.Color("8"),
		"steel":    lipgloss.Color("7"),
		"fairy":    lipgloss.Color("13"),
	}

	if color, ok := colors[pokemonType]; ok {
		return color
	}
	return lipgloss.Color("0")
}

func getShinyOrRegular(shiny bool) string {
	if shiny {
		return "shiny"
	}
	return "regular"
}

func getShinyOrRegularColor(shiny bool) lipgloss.Color {
	if shiny {
		return lipgloss.Color("11")
	}
	return lipgloss.Color("15")
}

func rollShiny(chance float64) bool {
	return rand.Float64() < chance
}

func formatPokemonInfo(dexBadge, name, genus, typeBadges, height, heightLabel, weight, weightLabel, flavorText string, mainColor lipgloss.Color) string {
	title := formatTitle(dexBadge, name, genus, mainColor)
	details := formatDetails(height, heightLabel, weight, weightLabel)
	flavorTextBox := formatFlavorText(flavorText)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		title,
		lipgloss.JoinVertical(
			lipgloss.Center,
			"",
			typeBadges,
			"",
			details,
			"",
			flavorTextBox,
		),
	)
}

func formatTitle(dexBadge, name, genus string, mainColor lipgloss.Color) string {
	return dexBadge + lipgloss.NewStyle().Bold(true).Foreground(mainColor).Render(fmt.Sprintf(" %s - %s", name, genus))
}

func formatDetails(height, heightLabel, weight, weightLabel string) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render(heightLabel+": "+lipgloss.NewStyle().Bold(true).Render(height)),
		lipgloss.NewStyle().Foreground(lipgloss.Color("15")).Render("	"+weightLabel+": "+lipgloss.NewStyle().Bold(true).Render(weight)),
	)
}

func formatFlavorText(flavorText string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("15")).
		Foreground(lipgloss.Color("15")).
		Padding(0, 1).
		Width(40).
		Render(flavorText)
}

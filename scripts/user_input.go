package GIFBot

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/jdkato/prose/v2"
)

// Liste des mots non importants
var WORST_W = map[string]struct{}{
	"je": {}, "suis": {}, "tu": {}, "il": {}, "elle": {}, "on": {}, "es": {},
	"a": {}, "de": {}, "plus": {}, "tqt": {}, "ai": {}, "est": {}, "sous": {},
	"vous": {}, "ils": {}, "elles": {}, "dehors": {}, "go": {}, "sommes": {},
	"étais": {}, "etais": {}, "ont": {}, "sont": {}, "êtes": {}, "etes": {},
	"dans": {}, "tres": {}, "super": {}, "bien": {}, "un": {}, "une": {},
	"le": {}, "la": {}, "les": {}, "quel": {}, "quelle": {}, "qu'on": {},
	"qu": {}, "'": {}, "": {}, "j'ai": {}, "mais": {}, "quoi": {}, "c'est": {},
	"te": {}, "faire": {}, "à": {}, "cest": {}, "aux": {}, "ce": {}, "gros": {}, "avec": {}, "meme": {}, "si": {}, "ça": {},
	"même": {}, "ces": {}, "cette": {}, "cela": {}, "cet": {}, "ne": {},
}

// Liste des thèmes communs et leurs synonymes
var themesCommuns = map[string][]string{
	"sports":             {"sports", "football", "basket", "soccer", "tennis", "baseball"},
	"films":              {"films", "cinéma", "action", "comédie", "drame", "romance", "horreur"},
	"musique":            {"musique", "rock", "pop", "jazz", "classique", "rap", "électro"},
	"météo":              {"météo", "pluie", "soleil", "neige", "tempête", "température", "prévisions"},
	"actualités":         {"actualités", "politique", "économie", "technologie", "science", "santé"},
	"nourriture":         {"nourriture", "cuisine", "recettes", "restaurants", "nutrition", "régime"},
	"voyages":            {"voyages", "vacances", "destination", "tourisme", "aventure", "séjour"},
	"jeux":               {"jeux", "jeux vidéo", "jeux de société", "puzzles", "stratégie", "rpg"},
	"livres":             {"livres", "littérature", "romans", "fiction", "non-fiction", "biographie"},
	"éducation":          {"éducation", "apprentissage", "études", "cours", "écoles", "universités"},
	"fitness":            {"fitness", "exercice", "entraînement", "yoga", "gym", "musculation"},
	"animaux":            {"animaux", "animaux de compagnie", "chiens", "chats", "faune", "zoo", "aquarium"},
	"mode":               {"mode", "vêtements", "style", "tendances", "design", "accessoires"},
	"art":                {"art", "peinture", "sculpture", "expositions", "artistes", "galeries"},
	"histoire":           {"histoire", "événements", "antique", "moderne", "civilisations", "guerres"},
	"relations":          {"relations", "amour", "rencontres", "mariage", "amis", "famille"},
	"technologie":        {"technologie", "gadgets", "logiciels", "matériel", "ordinateurs", "internet"},
	"automobiles":        {"automobiles", "voitures", "motos", "transport", "trains", "avions"},
	"maison":             {"maison", "intérieur", "décoration", "mobilier", "architecture", "rénovation"},
	"finance":            {"finance", "argent", "investissements", "banque", "cryptomonnaie", "budget"},
	"événements récents": {"événements récents", "tendances", "populaire", "viral", "célébrité", "divertissement"},
	"environnement":      {"environnement", "nature", "écologie", "conservation", "climat", "durabilité"},
	"politique":          {"politique", "gouvernement", "élections", "lois", "politiques", "international"},
}

func normalize(s string) string {
	var result strings.Builder
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsSpace(ch) {
			if unicode.IsLetter(ch) {
				result.WriteRune(unicode.ToLower(ch))
			} else {
				result.WriteRune(ch)
			}
		}
	}
	return result.String()
}

// Diviser la chaîne en mots importants en utilisant prose
func Split(s string) []string {
	var result []string
	doc, _ := prose.NewDocument(normalize(s))
	for _, tok := range doc.Tokens() {
		word := tok.Text
		if _, found := WORST_W[word]; !found && len(word) > 1 { // Ignore les mots courts
			result = append(result, word)
		}
	}
	return result
}

// Choisir le thème le plus pertinent
func ThemePicker(words []string) string {
	if len(words) == 0 {
		return "Nothing"
	} else if len(words) < 4 {
		s := ""
		for _, w := range words {
			s += w + " "
		}
		return s
	}

	// Compter la fréquence des mots et détecter les thèmes
	themeCount := make(map[string]int)
	for _, word := range words {
		for theme, synonyms := range themesCommuns {
			for _, synonym := range synonyms {
				if word == synonym {
					themeCount[theme]++
				}
			}
		}
	}

	var selectedTheme string
	maxCount := 0
	for theme, count := range themeCount {
		if count > maxCount || (count == maxCount && len(theme) > len(selectedTheme)) {
			selectedTheme = theme
			maxCount = count
		}
	}

	if selectedTheme == "" {
		s := ""
		for _, wr := range words {
			s += wr + " "
		}
		return s
	}
	return selectedTheme
}
func MessageToBot(s string) string {
	m := ThemePicker(Split(s))
	fmt.Println(m)
	return m
}

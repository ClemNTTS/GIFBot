package GIFBot

import "fmt"

// Split the following string in a []string with only important words and lower all words without ponctuation and number
func Split(s string) []string {
	t := []string{}
	save := 0
	inWord := false

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && inWord {
			if !isWorst(lower(s[save:i])) {
				t = append(t, lower(s[save:i]))
			}
			save = i + 1
			inWord = false
		} else {
			if !inWord && s[i] != ' ' {
				inWord = true
				save = i
			}
		}
	}

	t = append(t, lower(s[save:]))
	return t
}

func lower(s string) string {
	sf := ""
	for _, ch := range s {
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == 'é' || ch == 'è' || ch == 'ê' || ch == 'î' || ch == 'à' || ch == 'â' || ch == 'ô') {
			continue
		}
		if ch >= 'A' && ch <= 'Z' {
			sf += string(ch + 32)
		} else {
			sf += string(ch)
		}
	}
	return sf
}

var WORST_W []string = []string{"je", "suis", "tu", "il", "elle", "on", "es", "a", "de", "plus", "tqt", "ai", "est", "sous", "vous", "ils", "elles",
	"dehors", "go", "sommes", "étais", "etais", "ont", "sont", "êtes", "etes", "dans", "tres", "super", "bien", "un", "une", "le", "la", "les", "quel", "quelle", "qu'on", "qu", "'", "", "j'ai", "mais", "quoi", "c'est",
}

// if s is a non important word
func isWorst(s string) bool {
	for _, w := range WORST_W {
		if w == s {
			return true
		}
	}
	return false
}

// Pick the most important word of the sentence
func ThemePicker(t []string) string {
	if len(t) <= 0 {
		return "nothing"
	}
	tab := make(map[string]int)
	for _, w := range t {
		if !isWorst(w) {
			tab[w]++
		}
	}
	sav := ""
	max := 0
	for i, value := range tab {
		if value > max || (value == max && len(i) > len(sav)) {
			sav = i
			max = value
		}
	}

	return sav
}

func MessageToBot(s string) string {
	m := ThemePicker(Split(s))
	fmt.Println(m)
	return m
}

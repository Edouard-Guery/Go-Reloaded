package main

import (
	"testing"
)

func TestTraiterTexte(t *testing.T) {
	// Tableau des tests
	tests := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{"Conversion Hexadécimal", "1E (hex) files were added", "30 files were added"},
		{"Conversion Binaire", "It has been 10 (bin) years", "It has been 2 years"},
		{"Majuscule multiple", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"Règle a/an", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"Ponctuation simple", "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"Ponctuation de groupe", "I was thinking ... You were right", "I was thinking... You were right"},
		{"Guillemets", "As Elton John said: ' I am '", "As Elton John said: 'I am'"},
	}

	for _, test := range tests {
		t.Run(test.nom, func(t *testing.T) {
			resultat := TraiterTexte(test.entree)
			if resultat != test.attendu {
				t.Errorf("\nErreur sur le test : %s\nAttendu : %s\nObtenu  : %s", test.nom, test.attendu, resultat)
			}
		})
	}
}
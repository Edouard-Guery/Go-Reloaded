package main

import (
	"strings"
)

// FormaterPonctuation gère les espaces autour de la ponctuation 
func FormaterPonctuation(texte string) string {
	ponctuations := []string{".", ",", "!", "?", ":", ";"}

	// Enlever l'espace avant la ponctuation
	for _, p := range ponctuations {
		texte = strings.ReplaceAll(texte, " "+p, p)
	}

	//Forcer un espace après la ponctuation
	for _, p := range ponctuations {
		texte = strings.ReplaceAll(texte, p, p+" ")
	}

	texte = strings.ReplaceAll(texte, ". . . ", "...")
	texte = strings.ReplaceAll(texte, "? ! ", "?!")
	texte = strings.ReplaceAll(texte, "! ? ", "!?")
	texte = strings.ReplaceAll(texte, "! ! ", "!!")

	//Nettoyer les doubles espaces
	for strings.Contains(texte, "  ") {
		texte = strings.ReplaceAll(texte, "  ", " ")
	}

	//Gerer les guillemets simples
	parties := strings.Split(texte, "'")
	for i := 1; i < len(parties); i += 2 {
		parties[i] = strings.TrimSpace(parties[i])
	}
	texte = strings.Join(parties, "'")

	return strings.TrimSpace(texte)
}
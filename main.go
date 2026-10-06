package main

import (
	"fmt"
	"os"
)

func main() {
	// Verifie qu'on a bien 2 arguments après le nom du prog(entrée et sortie)
	if len(os.Args) != 3 {
		fmt.Println("Erreur : nombre d'arguments invalide ! Attendu: 2")
		return
	}
	fichierEntree := os.Args[1]
	fichierSortie := os.Args[2]

	//Lecture du fichier d'entrée
	donnees, err := os.ReadFile(fichierEntree)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Conversion des données en chaîne de caractères
	texteTexte := string(donnees)

	//Appel de la fonction principale de traitement
	texteModifie := TraiterTexte(texteTexte)

	// Ecriture du texte modifié dans le fichier de sortie
	err = os.WriteFile(fichierSortie, []byte(texteModifie), 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	}
}
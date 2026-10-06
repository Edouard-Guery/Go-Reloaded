package main

import (
	"strconv"
	"strings"
)

//TraiterTexte permet le  ligne par ligne 
func TraiterTexte(texte string) string {
	lignes := strings.Split(texte, "\n")
	var lignesModifiees []string

	for _, ligne := range lignes {
		lignesModifiees = append(lignesModifiees, TraiterLigne(ligne))
	}
	return strings.Join(lignesModifiees, "\n")
}

func TraiterLigne(ligne string) string {
	mots := strings.Fields(ligne)
	var resultat []string

	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		switch mot {
		case "(hex)":
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				val, _ := strconv.ParseInt(resultat[dernier], 16, 64)
				resultat[dernier] = strconv.FormatInt(val, 10)
			}
		case "(bin)":
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				val, _ := strconv.ParseInt(resultat[dernier], 2, 64)
				resultat[dernier] = strconv.FormatInt(val, 10)
			}
		case "(up)":
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = strings.ToUpper(resultat[dernier])
			}
		case "(low)":
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = strings.ToLower(resultat[dernier])
			}
		case "(cap)":
			if len(resultat) > 0 {
				dernier := len(resultat) - 1
				resultat[dernier] = Capitalize(resultat[dernier])
			}
		case "(up,", "(low,", "(cap,":
			if i+1 < len(mots) {
				nombreStr := strings.TrimRight(mots[i+1], ")")
				nombre, _ := strconv.Atoi(nombreStr)

				for j := 0; j < nombre; j++ {
					cible := len(resultat) - 1 - j
					if cible >= 0 {
						if mot == "(up," {
							resultat[cible] = strings.ToUpper(resultat[cible])
						} else if mot == "(low," {
							resultat[cible] = strings.ToLower(resultat[cible])
						} else if mot == "(cap," {
							resultat[cible] = Capitalize(resultat[cible])
						}
					}
				}
				i++ // passer le chiffre 
			}
		default:
			if mot == "a" || mot == "A" {
				if i+1 < len(mots) && CommenceParVoyelle(mots[i+1]) {
					mot += "n"
				}
			}
			resultat = append(resultat, mot)
		}
	}

	texteReconstruit := strings.Join(resultat, " ")
	return FormaterPonctuation(texteReconstruit)
}

func Capitalize(mot string) string {
	if len(mot) == 0 {
		return mot
	}
	return strings.ToUpper(string(mot[0])) + strings.ToLower(mot[1:])
}

func CommenceParVoyelle(mot string) bool {
	if len(mot) == 0 {
		return false
	}
	lettre := strings.ToLower(string(mot[0]))
	return strings.ContainsAny(lettre, "aeiouh")
}
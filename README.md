# 📝 Go-reloaded - Éditeur de texte CLI

Outil en ligne de commande (CLI) développé en Go (Golang) pour l'édition, la complétion et l'autocorrection automatisée de texte.

---

## 📖 Présentation
Réalisé dans le cadre de ma formation Bachelor Cyber à **YNOV Campus**, ce projet valide la manipulation de fichiers (`os`) et de chaînes de caractères (`strings`, `strconv`) en Go. 

Le but est de nettoyer un fichier texte brut contenant des balises spécifiques pour générer un fichier final formaté.
**Contrainte majeure :** Utilisation exclusive de la bibliothèque standard Go (aucun package externe).

---

## 🚀 Fonctionnalités
* **🔄 Conversions numériques :** Les balises `(hex)` et `(bin)` convertissent le mot précédent en base 10.
* **🔠 Casse des mots :** `(up)`, `(low)`, `(cap)` (majuscule, minuscule, capitalisée). Prise en charge des compteurs multiples (ex: `(up, 2)` modifie les 2 mots précédents).
* **✍️ Typographie :** Ajuste intelligemment l'espacement de la ponctuation (`. , ! ? : ;`) tout en préservant les groupes de ponctuation (`...`, `!?`).
* **💬 Guillemets :** Ferme les espaces à l'intérieur des guillemets simples (`'mot'`).
* **🧠 Grammaire :** Transforme automatiquement l'article `a` en `an` devant une voyelle ou un "h".

---

## 📁 Arborescence du Projet

```text
go-reloaded/
├── main.go               # Point d'entrée et gestion des fichiers (I/O)
├── formatage.go          # Moteur de traitement (parsing des mots, exécution des balises)
├── ponctuation.go        # Algorithmes de remplacement pour le nettoyage typographique
├── formatage_test.go     # Fichier de tests unitaires (couverture des règles)
├── go.mod                # Initialisation du module Go
├── sample.txt            # Fichier d'entrée (texte brut de test)
├── result.txt            # Fichier de sortie (généré par le programme)
└── README.md             # Documentation du projet
```

---

## 🛠️ Architecture Technique
Pour conserver un code propre, lisible et maintenable (Clean Code), le projet évite la sur-ingénierie et se découpe pragmatiquement en 3 fichiers de logique :
1. **L'Entrée/Sortie (`main.go`) :** Isole la logique système.
2. **Le Cerveau (`formatage.go`) :** Traite le texte ligne par ligne pour préserver la structure en paragraphes, et manipule les mots.
3. **Le Cosmétique (`ponctuation.go`) :** Intervient à la toute fin avec une logique de remplacements séquentiels robustes pour formater la typographie.

---

## 🕹️ Lancement rapide

### Prérequis
* **Go** : Version 1.22 ou supérieure.
* Un terminal standard.

### Exécuter le programme
Placez votre texte dans le fichier `sample.txt`, puis exécutez la commande à la racine du projet :

```bash
go run . sample.txt result.txt
```
*Le fichier `result.txt` sera automatiquement généré (ou écrasé) avec le texte corrigé.*

### Lancer l'audit de tests
Pour vérifier que toutes les règles fonctionnent correctement sans régression, lancez les tests unitaires :

```bash
go test -v
```